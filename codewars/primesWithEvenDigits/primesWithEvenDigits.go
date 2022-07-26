package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(F(1000))  // 887
	fmt.Println(F(1210))  // 1201
	fmt.Println(F(10000)) // 8887
	fmt.Println(F(500))   // 487
	fmt.Println(F(487))   // 467
	fmt.Println(F(6883))  // 6869
}

func F(n int) int {
	step := 1
	best := []int{0, 0}
	k := SieveOfEratosthenes(n)
	i := k[len(k)-step]
	for i >= 2 {
		step++
		evenDigits := countEvenDigits(i)
		max := maxEvenDigits(i)
		if best[0] < evenDigits {
			best = []int{evenDigits, i}
		}
		if max == best[0] {
			break
		}
		i = k[len(k)-step]
	}
	return best[1]
}

func countEvenDigits(n int) int {
	var evenDigits int
	for _, v := range strconv.Itoa(n) {
		if int(v)%2 == 0 {
			evenDigits++
		}
	}
	return evenDigits
}

func maxEvenDigits(n int) int {
	if len(strconv.Itoa(n)) == 1 {
		return 1
	}
	return len(strconv.Itoa(n)) - 1
}

func SieveOfEratosthenes(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] == true {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}
	var primes []int
	for p := 2; p < n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}
	return primes
}
