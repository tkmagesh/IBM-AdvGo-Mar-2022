package utils

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func printPrimes() {
	//generate primes
	primeNos := []int{}
	for no := 2; no <= 100; no++ {
		isPrime := true
		for i := 2; i < (no - 1); i++ {
			if no%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNos = append(primeNos, no)
		}
	}

	result := ""
	for _, primeNo := range primeNos {
		result += strconv.Itoa(primeNo) + ","
	}
	fmt.Println(result)
}

func utilsMain() {
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	printPrimes()
}
