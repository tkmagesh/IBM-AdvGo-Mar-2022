package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var c Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println("count = ", c.count)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	c.increment()
}
