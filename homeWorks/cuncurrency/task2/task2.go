package main

import "fmt"

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	var sum int
	ch := make(chan int)
	go func(ch chan int, n [][]int) {
		for _, val := range n {
			var res int
			for _, v := range val {
				res += v
			}
			ch <- res
		}
		close(ch)
	}(ch, n)
	for res := range ch {
		sum += res
	}
	fmt.Printf("result: %d", sum)
}
