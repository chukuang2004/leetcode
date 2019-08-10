package main

import (
	"fmt"
)

// https://leetcode-cn.com/explore/interview/card/bytedance/243/array-and-sorting/1017/
func bsearch(nums []int, start, end, target int) int {

	if start < 0 {
		return -1
	}
	if end >= len(nums) {
		return -1
	}
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	ret := -1
	if nums[mid] == target {
		ret = mid
	} else if nums[mid] > target {
		ret = bsearch(nums, start, mid-1, target)
	} else if nums[mid] < target {
		ret = bsearch(nums, mid+1, end, target)
	}

	return ret
}

func find(nums []int) int {

	size := len(nums)
	i, j := 0, size-1
	for i < j {
		if nums[i] < nums[i+1] {
			i++

		} else {
			return i
		}
		if nums[j-1] < nums[j] {
			j--
		} else {
			return j - 1
		}
	}

	return -1
}

func search(nums []int, target int) int {

	size := len(nums)
	if size == 0 {
		return -1
	}

	tmp := find(nums)
	ret := -1
	if tmp == -1 {
		ret = bsearch(nums, 0, size-1, target)
	} else {
		ret = bsearch(nums, 0, tmp, target)
		if ret == -1 {
			ret = bsearch(nums, tmp+1, size-1, target)
		}
	}
	return ret
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	// nums := []int{3, 1}
	// target := 1

	ret := search(nums, target)
	fmt.Println(ret)
}
