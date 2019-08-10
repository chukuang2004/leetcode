package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	tmpMap := make(map[string]int)

	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		j := i + 1
		m := len(nums) - 1
		tmp := 0 - nums[i]
		for j < m {

			if nums[j]+nums[m] > tmp {
				m--
			} else if nums[j]+nums[m] < tmp {
				j++
			} else if nums[j]+nums[m] == tmp {
				val := make([]int, 3)
				val[0] = nums[i]
				val[1] = nums[j]
				val[2] = nums[m]

				key := fmt.Sprintf("%d%d%d", val[0], val[1], val[2])
				_, ok := tmpMap[key]
				if !ok {
					ret = append(ret, val)
					tmpMap[key] = 1
				}

				j++
				m--
			}
		}
	}

	return ret
}

func main() {

	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}
