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
