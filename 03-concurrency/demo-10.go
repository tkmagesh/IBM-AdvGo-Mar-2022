package main

import (
	"fmt"
	"sync"
)

// DONT communicate by sharing memory
var result int

func main() {
	fmt.Println("main started")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println("Result = ", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
