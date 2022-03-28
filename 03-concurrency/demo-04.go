package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var count int
	fmt.Println(os.Args)
	if val, err := strconv.Atoi(os.Args[1]); err == nil {
		count = val
	} else {
		count = 1
	}
	fmt.Println(count)
	var wg sync.WaitGroup
	fmt.Println("main started")

	for i := 0; i < count; i++ {
		wg.Add(1)
		go f1(&wg) //scheduling
	}

	f2()
	wg.Wait()
	fmt.Println("main completed. Hit ENTER to exit")
	var input string
	fmt.Scanln(&input)
}

func f1(wg *sync.WaitGroup) {
	//fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
