package main

import (
	"fmt"
)

func isPalindrome(s string) bool {

	f := func(c byte) (byte, bool) {
		if c >= '0' && c <= '9' {
			return c, true
		}

		if c >= 'a' && c <= 'z' {
			return c, true
		}

		if c >= 'A' && c <= 'Z' {
			return c + ('a' - 'A'), true
		}

		return 0, false
	}

	for i, j := 0, len(s)-1; i < j; {

		a, ok := f(s[i])
		for ; i <= len(s)-1 && !ok;{
			a, ok = f(s[i])
            if !ok {
                i = i+1
            }
		}

		b, ok := f(s[j])
		for ; j >= 0 && !ok; {
			b, ok = f(s[j])
            if !ok {
                j = j -1
            }
		}

        if a != b {
            return false
        }

        if i<j{
            return false
        }

        if i==j

	}

}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
