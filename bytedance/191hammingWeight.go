package main

import (
	"fmt"
)

func hammingWeight1(num uint32) int {

	yu := num % 2
	ret := num / 2
	cnt := 0
	for ret > 0 {
		if yu == 1 {
			cnt = cnt + 1
		}

		yu = ret % 2
		ret = ret / 2
	}

	if yu == 1 {
		cnt = cnt + 1
	}

	return cnt
}

func hammingWeight2(num uint32) int {

	cnt := 0

	for num > 0 {

		if num&1 == 1 {
			cnt = cnt + 1
		}
		num = num >> 1
	}

	return cnt
}

func main() {

	fmt.Println(hammingWeight2(3))
}
