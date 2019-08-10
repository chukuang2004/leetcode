package main

import (
	"fmt"
)

func qsort1(src []int, start, end int) {

	if len(src) == 0 || start >= end {
		return
	}

	tmp := src[start]
	i := start + 1
	s, e := start, end

	for start < end {

		if src[i] < tmp {
			src[start], src[i] = src[i], src[start]
			start++

			i++
		} else {

			src[i], src[end] = src[end], src[i]

			end--
		}
	}

	qsort1(src, s, start-1)
	qsort1(src, start+1, e)
}

func qsort2(src []int, start, end int) {

}

func qsort3(src []int, start, end int) {

}

func main() {
	src := []int{1, 5, 3, 4, 2, 12, 123, 11, 450, 34, 45, 23, 37, 13, 14, 10}

	qsort1(src, 0, len(src)-1)

	fmt.Println(src)
}
