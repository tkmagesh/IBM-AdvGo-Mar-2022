package main

import "fmt"

/* func sumInts(list []int) int {
	result := 0
	for _, val := range list {
		result += val
	}
	return result
}

func sumFloats(list []float32) float32 {
	result := float32(0)
	for _, val := range list {
		result += val
	}
	return result
} */

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func sum[T Numbers](list []T) T {
	var result T
	for _, val := range list {
		result += val
	}
	return result
}

func filter[T Numbers](list []T, predicate func(T) bool) []T {
	var result []T
	for _, val := range list {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	ints := []int{3, 1, 4, 2, 5, 8, 6, 7, 9}
	//intSum := sumInts(ints)
	//intSum := sum[int](ints)
	intSum := sum(ints)
	fmt.Println(intSum)

	evenNos := filter(ints, func(val int) bool {
		return val%2 == 0
	})
	fmt.Println(evenNos)

	floats := []float32{3, 1, 4.5, 2.5, 5, 8, 6, 7, 9}
	//floatSum := sumFloats(floats)
	//floatSum := sum[float32](floats)
	floatSum := sum(floats)
	fmt.Println(floatSum)
}
