package main

import "fmt"

type Employee struct {
	salary float32
}

func incrementSalary(e *Employee) {
	e.salary = e.salary + (e.salary * 0.1)
}

func main() {
	e := Employee{10000}
	fmt.Println(e)
	fmt.Println("After incrment:")
	incrementSalary(&e)
	fmt.Println(e)
}
