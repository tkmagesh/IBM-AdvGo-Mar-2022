package main

import (
	"fmt"
	"time"
	"worker-demo/worker"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"Magesh",
	"Ganesh",
	"Ramesh",
	"Rajesh",
	"Suresh",
}

type NamePrinter struct {
	name  string
	delay time.Duration
}

/* Task = any object implementing the Task() function */
func (np *NamePrinter) Task() {
	fmt.Println("Task execution commenced")
	time.Sleep(np.delay)
	fmt.Println("Name Printer - Name : ", np.name)
}

func main() {
	w := worker.New(5)
	timerCounter := 1
	for idx := 0; idx < 2; idx++ {
		for _, name := range names {
			np := &NamePrinter{
				name:  name,
				delay: time.Duration(time.Duration(timerCounter) * time.Second),
			}
			timerCounter++
			w.Run(np)
		}
	}
	fmt.Println("All tasks are assigned")
	w.Shutdown()
}
