package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	size := len(digits)

	over := true
	for i := size - 1; i >= 0; i-- {
		val := digits[i]
		if val == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			over = false
			break
		}
	}

	ret := make([]int, 1)
	ret[0] = 1
	if over {
		digits = append(ret, digits...)
	}

	return digits
}

func main() {

	ret := plusOne([]int{1, 2, 4})

	fmt.Println(ret)
}
