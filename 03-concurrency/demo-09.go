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
		go func(id int) {
			defer wg.Done()
			//count++
			mutex.Lock()
			{
				count = count + 1
			}
			mutex.Unlock()
			fmt.Println("id = ", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("count = ", count)
}
