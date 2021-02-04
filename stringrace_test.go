package go_learning

import (
	"testing"
	"time"
)

var a = "0"

func TestStringRace(t *testing.T) {
	ch := make(chan string)
	go func() {
		i := 0
		for {
			if i%2 == 0 {
				a = "aa"
			} else {
				a = "0"
			}
			i++
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < (1 << 31); i++ {
			b := a
			if b != "0" && b != "aa" {
				ch <- b
			}
		}
	}()

	tick := time.After(5 * time.Second)
	stopLoop := false
	for {
		select {
		case c := <-ch:
			t.Log(c)
		case stopTime := <-tick:
			t.Log(stopTime)
			stopLoop = true
		}
		if stopLoop {
			break
		}
	}
}

func TestSelect(t *testing.T) {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
		ch2 <- i
		ch3 <- i
	}
	var n int
	for {
		select {
		case n = <-ch1:
			t.Log("ch1", n)
		default:
			select {
			case n = <-ch2:
				t.Log("ch2", n)
			default:
				select {
				case n = <-ch3:
					t.Log("ch3", n)
				default:
					return
				}
			}
		}
	}
}
