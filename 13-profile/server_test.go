package main

import "testing"

func Benchmark_isPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(97)
	}
}

func Benchmark_getPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimes()
	}
}
