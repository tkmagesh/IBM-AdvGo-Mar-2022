package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go doSomething(wg, ctx)
	go func() {
		fmt.Println("Hit ENTER to stop....")
		var input string
		fmt.Scanln(&input)
		cancel()
	}()
	wg.Wait()

}

func doSomething(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("doing something")
		}
	}
}
