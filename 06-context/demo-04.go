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

	ctxWithValue := context.WithValue(context.Background(), "root-key", "root-value")
	ctx, cancel := context.WithCancel(ctxWithValue)
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
	fmt.Println("[doSomething] Value from context = ", ctx.Value("root-key"))
	wg.Add(1)
	dsCtx := context.WithValue(ctx, "child-key", "child-value")
	timeoutCtx, cancel := context.WithTimeout(dsCtx, 10*time.Second)
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
	fmt.Println("[doSomethingElse] Value from context[root] = ", ctx.Value("root-key"))
	fmt.Println("[doSomethingElse] Value from context[child] = ", ctx.Value("child-key"))
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
