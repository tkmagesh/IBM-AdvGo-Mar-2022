package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	primeCh := genPrimes(count)
	for i := 0; i < 11; i++ {
		fmt.Println(<-primeCh)
	}
	fmt.Println("Done")
}

func genPrimes(count int) <-chan int {
	ch := make(chan int)
	go func() {
		var no = 2
		for count > 0 {
			if isPrime(no) {
				ch <- no
				count--
				time.Sleep(500 * time.Millisecond)
			}
			no++
		}
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
