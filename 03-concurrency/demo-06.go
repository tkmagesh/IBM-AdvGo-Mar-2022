package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

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
	//count++
	mutex.Lock()
	{
		count = count + 1
	}
	mutex.Unlock()
}
