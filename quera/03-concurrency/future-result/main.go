package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
}

type Task func() string

func Async(t Task) *FutureResult {
	fr := &FutureResult{
		Done:       atomic.Bool{},
		ResultChan: make(chan string, 1),
	}
	go func() {
		fr.ResultChan <- t()
		fr.Done.Store(true)
	}()
	return fr
}

func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	fr := &FutureResult{
		Done:       atomic.Bool{},
		ResultChan: make(chan string, 1),
	}

	taskChan := make(chan string, 1)
	taskRunner := func(ch chan string) {
		ch <- t()
	}

	go func() {
		go taskRunner(taskChan)
		select {
		case res := <-taskChan:
			fr.ResultChan <- res
			fr.Done.Store(true)
		case <-time.After(timeout + time.Millisecond):
			fr.ResultChan <- "timeout"
			fr.Done.Store(false)
		}
	}()
	return fr
}

func (fResult *FutureResult) Await() string {
	return <-fResult.ResultChan
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	combined := &FutureResult{
		Done:       atomic.Bool{},
		ResultChan: make(chan string, len(fResults)),
	}

	go func() {
		for _, single := range fResults {
			combined.ResultChan <- single.Await()
		}
		combined.Done.Store(true)
	}()

	return combined
}

//func CombineFutureResultsOrderLess(fResults ...*FutureResult) *FutureResult {
//	combined := &FutureResult{
//		Done:       atomic.Bool{},
//		ResultChan: make(chan string, len(fResults)),
//	}
//	var wg sync.WaitGroup
//	for _, single := range fResults {
//		if single.Done.Load() {
//			combined.ResultChan <- single.Await()
//		}
//		wg.Add(1)
//		go func(single *FutureResult) {
//			defer wg.Done()
//			combined.ResultChan <- single.Await()
//		}(single)
//	}
//	go func() {
//		wg.Wait()
//		combined.Done.Store(true)
//	}()
//	return combined
//}
