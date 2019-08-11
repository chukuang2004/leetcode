package main

import (
	"fmt"
)

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates(nums []int) int {

	i, j := 0, 0
	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	} else {
		j = 1
	}

	for ; j < len(nums); j++ {
		if nums[i] == nums[j] {

		} else {
			i++
			if (j - i) > 0 {
				nums[i] = nums[j]
			}
		}
	}

	return i + 1
}

func removeDuplicates1(nums []int) int {

	i, j := 0, 0
	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	} else {
		j = 1
	}

	flag := false
	for ; j < len(nums); j++ {
		if nums[i] == nums[j] {
			if !flag {
				if nums[i] != nums[i+1] {
					nums[i+1] = nums[j]
				}
				i++
			}
			flag = true
		} else {

			flag = false
			i++
			if (j - i) > 0 {
				nums[i] = nums[j]
			}
		}
	}

	return i + 1
}

func main() {
	fmt.Println(removeDuplicates1([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
