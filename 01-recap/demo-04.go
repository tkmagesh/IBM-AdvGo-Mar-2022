/* Higher order functions */
package main

import "fmt"

func main() {
	//anonymous functions
	func() {
		fmt.Println("Anonymous function invoked")
	}()
	fmt.Println("anonymous function invocation completed")

	func(x, y int) {
		fmt.Println("Anonymous function invoked, result = ", x+y)
	}(100, 200)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("Result from anonymous function = ", result)

	/* functions assigned to variables */
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))
}
