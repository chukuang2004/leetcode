package main

import (
	"fmt"
)

// nums 有序
func twoSum(nums []int, target int) []int {

	if len(nums) <= 1 || nums[1] > target {
		return nil
	}

	start := 0
	end := -1

	for i, v := range nums {
		if v <= target {
			end = i
		} else {
			break
		}
	}

	tmp := target / 2
	middle := 0
	for i, v := range nums {

		if v > tmp {
			middle = i
			break
		} else if v == tmp {
			return nil
		}
	}

	ret := []int{}
	for i := start; i < middle; i++ {
		tmp := target - nums[i]

		for j := middle; j <= end; j++ {

			if tmp == nums[j] {
				ret = append(ret, nums[i])
				ret = append(ret, nums[j])

			}
		}
	}

	return ret
}

func main() {
	array := []int{2, 7, 11, 15}
	target := 9

	ret := twoSum(array, target)

	fmt.Println(ret)
}
