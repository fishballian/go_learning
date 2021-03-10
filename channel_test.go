package go_learning

import (
	"strconv"
	"testing"
	"time"
)

func TestReadWriteChan(t *testing.T) {
	rw := NewReadWriteChan()
	for i := 0; i < 10000; i++ {
		n := i
		go func() {
			rw.Set("seq:" + strconv.Itoa(n))
			t.Log(rw.Get())
		}()
	}
	time.Sleep(time.Second)
	t.Log(rw.Get())
}

func TestSelectSend(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	select {
	case ch <- 1:
		t.Log("send")
	default:
		t.Log("default")
	}
}
