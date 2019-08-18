package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	s := len(haystack)

	for i := 0; i < s; i++ {
		if haystack[i] == needle[0] {
			ss := len(needle)
			if i+ss-1 >= s {
				return -1
			}

			flag := true
			for j := 1; j < ss; j++ {
				if haystack[i+j] != needle[j] {
					flag = false
					break
				}
			}

			if flag {
				return i
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("asaa", "aa"))
}
