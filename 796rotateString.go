package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {

	alen := len(A)
	blen := len(B)

	if alen != blen {
		return false
	}

	if alen == 0 {
		return true
	}

	if alen == 1 && A[0] == B[0] {
		return true
	}

	for i := alen - 1; i > 0; i-- {

		if A[i] == B[0] && A[alen-1] == B[alen-i-1] {
			if A[i-1] == B[alen-1] && A[0] == B[alen-i] {
				if 0 == strings.Compare(A[0:i], B[alen-i:]) &&
					0 == strings.Compare(A[i:], B[0:alen-i]) {
					return true
				}
			}
		}
	}

	return false
}

func main() {

	fmt.Println(rotateString("abcdf", "cdeab"))
}
