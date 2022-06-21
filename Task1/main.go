package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	uniqElements := make(map[int]bool)
	res := make([]int, 0)

	for _, v := range arr {
		ok := uniqElements[v]
		if !ok {
			res = append(res, v)
			uniqElements[v] = true
		}
	}

	fmt.Println(res)
}
