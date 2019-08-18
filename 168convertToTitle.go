package main

import (
	"fmt"
)

func convertToTitle(n int) string {

	if n < 1 {
		return ""
	}

	left := n / 26
	yu := n % 26

	if left > 0 && yu == 0 {
		return convertToTitle(left-1) + "Z"
	}

	return convertToTitle(left) + fmt.Sprintf("%c", yu-1+'A')
}

func main() {
	fmt.Printf(convertToTitle(701))
}
