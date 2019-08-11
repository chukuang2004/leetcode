package main

import (
	"fmt"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	y := func(v int) int {
		yu := v % 10
		tmp := v / 10
		ret := yu
		for tmp > 0 {
			yu = tmp % 10
			ret = ret*10 + yu
			tmp = tmp / 10
		}

		return ret
	}(x)

	if x == y {
		return true
	}

	return false
}

func main() {

}
