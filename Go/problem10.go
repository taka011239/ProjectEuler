package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func genPrime(limit int, primes chan<- int) {
	primeFlgs := make([]bool, limit)
	for i := 0; i < len(primeFlgs); i++ {
		primeFlgs[i] = true
	}

	primes <- 2

	for i := 3; i < limit; i += 2 {
		if primeFlgs[i] == true {
			primes <- i
		}

		for j := i; j < limit; j += i {
			primeFlgs[j] = false
		}
	}

	close(primes)
}

func calcTotal(primes <-chan int, total chan<- int) {
	sum := 0
	for p := range primes {
		sum += p
	}
	total <- sum
}

func main() {
	runtime.GOMAXPROCS(2)

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: integer")
		return
	}

	primes := make(chan int, 100)
	total := make(chan int)

	go calcTotal(primes, total)
	go genPrime(limit, primes)

	fmt.Printf("%d\n", <-total)
	close(total)
}
