package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() //scheduling
	f2()
	time.Sleep(1 * time.Millisecond) //DO NOT DO THIS

	//DO NOT DO THIS
	/*
		var input string
		fmt.Scanln(&input)
	*/
	fmt.Println("main completed")
	panic("test")
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
