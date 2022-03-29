package main

import (
	"fmt"
	"sync"
)

func main() {
	//using waitgroup

	/*
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go fnWg(wg)
		}
		wg.Wait()
	*/

	//using channels
	done := make(chan any)
	for i := 0; i < 10; i++ {
		go fnCh(done)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println("done")
}

func fnWg(wg *sync.WaitGroup) {
	fmt.Println("fn invoked")
	wg.Done()
}

func fnCh(doneCh chan any) {
	fmt.Println("fn invoked")
	doneCh <- struct{}{}
}
