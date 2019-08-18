package main

import (
	"fmt"
)

func powerfulIntegers(x int, y int, bound int) []int {
	xx := 1
	tmp := make(map[int]struct{})
	ret := []int{}
	for xx < bound {
		yy := 1
		left := bound - xx
		for yy <= left {
			tmp[xx+yy] = struct{}{}
			yy = yy * y
			if yy == 1 {
				break
			}
		}
		xx = xx * x
		if xx == 1 {
			break
		}
	}

	for k, _ := range tmp {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	fmt.Println(powerfulIntegers(1, 2, 10))
}
