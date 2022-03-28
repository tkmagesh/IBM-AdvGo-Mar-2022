package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {
	var x any
	x = 100
	//x = "this is a string"
	//x = true
	x = struct{}{}
	x = Employee{100, "Magesh", "Kuppan"}
	//x = 10.25
	if val, ok := x.(int); ok {
		fmt.Println(val + 100)
	} else {
		fmt.Println("x is not an int")
	}

	if emp, ok := x.(Employee); ok {
		fmt.Println("Employee = ", emp)
	} else {
		fmt.Println("x is not an employee")
	}

	switch val := x.(type) {
	case int:
		fmt.Println(val + 100)
	case string:
		fmt.Println("x is string with length : ", len(val))
	case Employee:
		fmt.Println("Employee = ", val)
	default:
		fmt.Println("unknown type")
	}

}
