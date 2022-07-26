package main

import (
	"fmt"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	var sum int
	ch := make(chan int, len(n))

	for _, val := range n {
		go func(val []int, ch chan int) {
			var res int
			for _, v := range val {
				res += v
			}
			ch <- res
		}(val, ch)
	}

	for i := 0; i < len(n); i++ {
		sum += <-ch
	}
	close(ch)
	fmt.Printf("result: %d\n", sum)
}
