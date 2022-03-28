package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")

	ch := make(chan int)
	go add(100, 200, ch)
	time.Sleep(2 * time.Second)
	result := <-ch
	fmt.Println("Result = ", result)

	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result

}
