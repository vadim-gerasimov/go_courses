package main

import "fmt"

func DigitalRoot(n int) int {
	return (n-1)%9 + 1
}

func main() {
	fmt.Println(DigitalRoot(16))     // 7
	fmt.Println(DigitalRoot(195))    // 6
	fmt.Println(DigitalRoot(992))    // 2
	fmt.Println(DigitalRoot(167346)) // 9
	fmt.Println(DigitalRoot(0))      // 0
}
