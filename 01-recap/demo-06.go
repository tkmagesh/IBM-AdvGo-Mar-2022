package main

import "fmt"

func main() {

	/*
		add(100,200)
		subtract(100,200)
	*/

	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	logAdd := getLogOperation(add)
	logSubtract := getLogOperation(subtract)

	logAdd(100, 200)
	logSubtract(100, 200)

}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("operation started")
		oper(x, y)
		fmt.Println("operation completed")
	}
}

func add(x, y int) {
	fmt.Println("Add Result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result = ", x+y)
}
