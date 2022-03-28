package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("count = ", count)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&count, 1)
}
