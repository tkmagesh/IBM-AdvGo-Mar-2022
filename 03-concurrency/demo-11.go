package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main started")
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	result := <-ch
	wg.Wait()
	fmt.Println("Result = ", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
