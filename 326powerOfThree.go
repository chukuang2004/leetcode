package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	ret := n

	for {
		yu := ret % 3
		ret = ret / 3

		if ret == 0 {
			if yu == 1 {
				return true
			} else {
				return false
			}
		} else if yu != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPowerOfThree(6))
}
