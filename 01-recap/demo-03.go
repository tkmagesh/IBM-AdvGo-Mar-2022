package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                         //=> 100
	fmt.Println(sum(10))                                     //=> 10
	fmt.Println(sum())                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                  //=> 100
	fmt.Println(sum(10, 20, []any{30, 40, []int{10, 20}}))   //=> 130
	fmt.Println(sum(10, 20, []any{30, 40, []any{10, "20"}})) //=> 130
}

func sum(data ...any) int {
	result := 0
	for _, v := range data {
		switch val := v.(type) {
		case int:
			result += val
		case string:
			if no, ok := strconv.Atoi(val); ok == nil {
				result += no
			}
		case []any:
			result += sum(val...)
		case []int:
			intfList := make([]any, len(val))
			for i, v := range val {
				intfList[i] = v
			}
			result += sum(intfList...)
		}
	}
	return result
}
