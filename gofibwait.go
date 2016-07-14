package gofibwait

import (
	"fmt"
	"time"
)

type Waiter struct {
	max     int
	wait    int
	waitFib int
}

func NewWaiter(max int) Waiter {
	return Waiter{wait: 0, waitFib: 1, max: max}
}

func (w *Waiter) Wait() time.Duration {
	dur := time.Duration(w.wait) * time.Second
	fmt.Printf("sleeping %s\n", dur)
	time.Sleep(dur)
	if w.wait < w.max {
		fib := w.waitFib
		w.waitFib = w.wait
		w.wait = w.wait + fib
	}
	return dur
}

func (w *Waiter) Reset() {
	w.wait = 0
	w.waitFib = 1
}
