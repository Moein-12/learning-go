package main

import (
	"fmt"
	"sync"
)


func markNonPrime(lcd int, n int, wg *sync.WaitGroup, primes []bool) {
	for i := lcd; i < n; i++ {
		if primes[i] && i % lcd != 0 {
			wg.Add(1)
			go markNonPrime(i, n, wg, primes)
			break
		}
	} 
	
	for i := lcd * 2; i < n; i += lcd {
		primes[i] = false
	}
	wg.Done()
}


func sieve(n int) []int {
	lst := make([]bool, n + 1)
	var wg sync.WaitGroup
	
	for i := 2; i < n; i++ {
		lst[i] = true
	}

	wg.Add(1)
	markNonPrime(2, n, &wg, lst)
	wg.Wait()
	
	
	primes := make([]int, 0)

	for i := 0; i < n + 1; i++ {
		if lst[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	fmt.Println(sieve(50))
}
