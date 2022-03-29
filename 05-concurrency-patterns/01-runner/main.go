package main

import (
	"fmt"
	"os"
	"runner-demo/runner"
	"time"
)

func main() {
	/*
		initialize the runner with a timeout
		Add multiple tasks to the runner
		Start the runner
		if all the tasks are completed within the given time, report "success"
		if the tasks are not completed within the given time, report "timeout"
		exit if the execution is interrupted by an OS interrupt
	*/

	fmt.Printf("Process %d started....\n", os.Getpid())
	//initialize the runner with a timeout
	timeout := 15 * time.Second
	r := runner.New(timeout)

	//Add multiple tasks to the runner
	r.Add(createTask(3))
	r.Add(createTask(4))
	r.Add(createTask(7))

	//Start the runner
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("tasks timed out")
		default:
			fmt.Println("unknown error : ", err)
		}
	} else {
		fmt.Println("success")
	}
	//if all the tasks are completed within the given time, report "success"
	//if the tasks are not completed within the given time, report "timeout"
	//exit if the execution is interrupted by an OS interrupt ( kill -2 <pid> )

}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Printf("Processing Task #%d\n", id)
		time.Sleep(time.Duration(t) * time.Second)
	}
}
