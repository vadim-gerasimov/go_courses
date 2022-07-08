package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "-1 3 4 78 90000 -1234235 77839204 -9999999 2 75 23 -3 -4 756"
	var res string
	var min, max int32

	sf := strings.Fields(input)

	for i, val := range sf {
		intVal, _ := strconv.Atoi(val)
		if i == 0 {
			max = int32(intVal)
			min = int32(intVal)
		}
		if int32(intVal) < min {
			min = int32(intVal)
		}
		if int32(intVal) > max {
			max = int32(intVal)
		}
	}

	res = fmt.Sprintf("%d %d", max, min)
	fmt.Println(res)

}
