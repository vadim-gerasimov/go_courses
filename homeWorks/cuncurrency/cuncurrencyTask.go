package main

import (
	"fmt"
	"sync"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	var wg sync.WaitGroup
	for i, v := range n {
		wg.Add(1)
		go func(v []int, i int) {
			defer wg.Done()
			sum(v, i)
		}(v, i)
	}
	wg.Wait()
}

func sum(v []int, i int) {
	var sum int
	for _, val := range v {
		sum += val
	}
	fmt.Printf("slice %d: %d\n", i, sum)
}
