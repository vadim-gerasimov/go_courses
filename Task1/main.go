package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	uniqElements := make(map[int]bool)
	res := make([]int, 0)

	for _, val := range arr {
		uniqElements[val] = true
	}

	for num, _ := range uniqElements {
		res = append(res, num)
	}

	fmt.Println(res)
}
