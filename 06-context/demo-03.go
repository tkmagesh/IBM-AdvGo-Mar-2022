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
	ctx, cancel := context.WithCancel(context.Background())
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
	wg.Add(1)
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	go doSomethingElse(wg, timeoutCtx)
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

func doSomethingElse(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("doing something else")
		}
	}
	fmt.Println("doSomethingElse completed")
}
