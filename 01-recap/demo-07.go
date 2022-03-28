package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment())
	fmt.Println(increment())
	doSomething()
	fmt.Println(increment())
	fmt.Println(increment())
}

func getIncrement() func() int {
	var count int //closure
	increment := func() int {
		count++
		return count
	}
	return increment
}

func doSomething() {
	//count = 1000
}
