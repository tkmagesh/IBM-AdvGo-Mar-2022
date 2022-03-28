package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	stop := time.After(20 * time.Second)
LOOP:
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-stop:
			fmt.Println("Done")
			break LOOP
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
