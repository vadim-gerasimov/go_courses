package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(Partitions(1))  // 1
	fmt.Println(Partitions(3))  // 2
	fmt.Println(Partitions(5))  // 3
	fmt.Println(Partitions(10)) // 10
	fmt.Println(Partitions(20)) //64
}

func Partitions(n int) *big.Int {
	p := make([]*big.Int, n+1)
	p[0] = big.NewInt(1)
	for i := 1; i <= n; i++ {
		p[i] = big.NewInt(0)
	}
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			p[j] = new(big.Int).Add(p[j], p[j-i])
		}
	}
	return p[n]
}
