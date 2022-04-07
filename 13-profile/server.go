package main

import (
	"bytes"
	"math"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
)

func isPrime(no int) bool {
	//max := no - 1
	//max := no / 2
	max := int(math.Sqrt(float64(no)))
	for i := 2; i <= max; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func getPrimes() []int {
	//primeNos := []int{}
	primeNos := make([]int, 40)
	for no := 2; no <= 100; no++ {
		if isPrime(no) {
			primeNos = append(primeNos, no)
		}
	}
	return primeNos
}

/*
func preparePrimes() string {
	primeNos := getPrimes()
	result := ""
	for _, primeNo := range primeNos {
		result += strconv.Itoa(primeNo) + ","
	}
	return result
}
*/

//pool is recycled for every GC

var pool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func preparePrimes(primeNos *[]int) *[]byte {
	result := pool.Get().(*bytes.Buffer)
	for _, primeNo := range *primeNos {
		result.Write([]byte(strconv.Itoa(primeNo)))
		result.WriteRune(',')
	}
	resultBytes := result.Bytes()
	result.Reset()
	pool.Put(result)
	return &resultBytes
}

func primeHandler(w http.ResponseWriter, r *http.Request) {
	//generate primes
	primeNos := getPrimes()
	result := preparePrimes(&primeNos)
	w.Write(*result)
}

func main() {
	http.HandleFunc("/primes", primeHandler)
	http.ListenAndServe(":8080", nil)
}
