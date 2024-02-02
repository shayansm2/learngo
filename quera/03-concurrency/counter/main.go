package main

import (
	"fmt"
	"sync/atomic"
)

type Counter int32

func (c *Counter) Increment() {
	for {
		old := *c
		new := old + 1
		if atomic.CompareAndSwapInt32((*int32)(c), int32(old), int32(new)) {
			break
		}
	}
}

func (c *Counter) Value() int32 {
	return int32(atomic.LoadInt32((*int32)(c)))
}

func main() {
	var counter Counter

	counter.Increment()
	fmt.Println("Counter Value:", counter.Value())
}
