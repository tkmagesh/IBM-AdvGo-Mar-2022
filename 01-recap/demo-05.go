package main

import "fmt"

func main() {

	/*
		add(100,200)
		subtract(100,200)
	*/

	/*
		fmt.Println("operation started")
		add(100,200)
		fmt.Println("operation completed")

		fmt.Println("opeartion started")
		subtract(100,200)
		fmt.Println("operation completed")
	*/

	/*
		logAddOperation(100, 200)
		logSubtractOperation(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)

}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("operation started")
	oper(x, y)
	fmt.Println("operation completed")
}

func logAddOperation(x, y int) {
	fmt.Println("operation started")
	add(x, y)
	fmt.Println("operation completed")
}

func logSubtractOperation(x, y int) {
	fmt.Println("opeartion started")
	subtract(x, y)
	fmt.Println("operation completed")
}

func add(x, y int) {
	fmt.Println("Add Result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result = ", x+y)
}
