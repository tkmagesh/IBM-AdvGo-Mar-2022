package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("main started")
	//wg.Add(100)
	wg.Add(1)
	go f1(&wg) //scheduling
	f2()
	wg.Wait()
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
