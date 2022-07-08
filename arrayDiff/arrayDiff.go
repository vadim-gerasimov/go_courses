package main

import "fmt"

func ArrayDiff(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range b {
		m[item] = true
	}
	diff := []int{}
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}
func main() {
	fmt.Println(ArrayDiff([]int{1, 2}, []int{1}))       //1
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{1}))    //2,2
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{2}))    // 1
	fmt.Println(ArrayDiff([]int{1, 2, 2}, []int{}))     // 1,2,2
	fmt.Println(ArrayDiff([]int{}, []int{1, 2}))        //
	fmt.Println(ArrayDiff([]int{1, 2, 3}, []int{1, 2})) // 3
}
