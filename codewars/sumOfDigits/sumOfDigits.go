package main

import "fmt"

func DigitalRoot(n int) int {
	sum := 0
	for {
		sum += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	if sum > 9 {
		sum = DigitalRoot(sum)
	}
	return sum
}

func main() {
	fmt.Println(DigitalRoot(16))     // 7
	fmt.Println(DigitalRoot(195))    // 6
	fmt.Println(DigitalRoot(992))    // 2
	fmt.Println(DigitalRoot(167346)) // 9
	fmt.Println(DigitalRoot(0))      // 0
}
