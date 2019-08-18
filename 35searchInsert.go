package main

import (
	"fmt"
)

//迭代8ms  遍历4ms
func searchInsert(nums []int, target int) int {

	s := len(nums)
	if s < 1 {
		return 0
	}
	//    else if s == 1 {
	// 	if nums[0] > target {
	// 		return 0
	// 	} else if nums[0] == target {
	// 		return 0
	// 	} else {
	// 		return 1
	// 	}
	// }
	// for {
	// 	mid := s / 2

	// 	if nums[mid] > target {
	// 		return searchInsert(nums[:mid], target)
	// 	} else if nums[mid] == target {
	// 		return mid
	// 	} else {
	// 		return mid + searchInsert(nums[mid:], target)
	// 	}
	// }

	low := 0
	if nums[low] == target {
		return low
	}
	high := s - 1
	if nums[high] == target {
		return high
	}

	mid := (high + low) / 2
	for low <= high {

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] == target {
			return mid
		} else {
			low = mid + 1
		}
		mid = (high + low) / 2
	}

	return low
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 4, 5, 6}, 2))
}
