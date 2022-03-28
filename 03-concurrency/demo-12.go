package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Attempting to write")
		ch <- 100
		fmt.Println("write successful")
	}()
	fmt.Println("Attempting to read")
	data := <-ch
	fmt.Println("read successful")
	fmt.Println(data)
}
