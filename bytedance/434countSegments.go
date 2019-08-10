package main

import (
	"fmt"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := 0
	start := false
	for ; i < len(s); i++ {
		if !start {
			if s[i] != ' ' {
				start = true
			}
		} else {
			if s[i] == ' ' {
				break
			}
		}
	}

	if !start {
		return 0
	}

	if i < len(s) {
		return 1 + countSegments(s[i+1:])
	}

	return 1
}

func main() {
	fmt.Println(countSegments("  he llo, .  world"))
}
