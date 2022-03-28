package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan any)
	primeCh := genPrimes(stopCh)

	go func() {
		var input string
		fmt.Scanln(&input)
		//stopCh <- true
		close(stopCh)
	}()

	for primeNo := range primeCh {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(stopCh chan any) <-chan int {
	ch := make(chan int)
	go func() {
		no := 2
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				if isPrime(no) {
					ch <- no
					time.Sleep(500 * time.Millisecond)
				}
				no++
				continue LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/* func timeOut(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
