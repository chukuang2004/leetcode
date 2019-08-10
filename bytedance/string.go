package main

import (
	"bytes"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	max := 0

	if len(s) == 1 {
		return 1
	}

	src := bytes.NewBufferString(s).Bytes()

	for start := 0; start < len(src)-1; start++ {

		bucket := make([]int, 128)

		index := src[start] - 0
		bucket[index] = 1
		end := start + 1

		for ; end < len(src); end++ {
			index := src[end] - 0
			if bucket[index] != 0 {
				break
			}
			bucket[index] = 1
		}

		tmp := end - start

		if max < tmp {
			max = tmp
		}
	}

	return max
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	buf := make([][]byte, len(strs))

	size := len(strs[0])
	for i := 0; i < len(strs); i++ {
		buf[i] = bytes.NewBufferString(strs[i]).Bytes()

		if len(strs[i]) < size {
			size = len(strs[i])
		}
	}

	end := 0
	for ; end < size; end++ {
		tmp := buf[0][end]
		get := false
		for i := 0; i < len(strs); i++ {
			if tmp != buf[i][end] {
				get = true
				break
			}
		}

		if get {
			break
		}
	}

	return strs[0][:end]
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}

	b1 := bytes.NewBufferString(s1).Bytes()
	b2 := bytes.NewBufferString(s2).Bytes()

	for i := 0; i <= len(s2)-len(s1); i++ {

		for j := 0; j < len(s1); j++ {

			get := false
			for m := 0; m < len(s1); m++ {
				if b1[j] != b2[i+j+m] {
					get = true
					break
				}
			}

			if !get {
				return true
			}
		}
	}

	return false
}

func restoreIpAddresses(s string) []string {

	return nil
}

func main() {

	cnt := lengthOfLongestSubstring("ppwwfasg")
	fmt.Println(cnt)

	strs := []string{"sdhgqwqe", "sdhg"}
	str := longestCommonPrefix(strs)
	fmt.Println(str)

	flag := checkInclusion("ab", "eidbaooo")
	fmt.Println(flag)
}
