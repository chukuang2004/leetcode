package main

import (
	"fmt"
)

// 动态规划 记忆化搜索递归
func tribonacci(n int) int {

	mem := make([]int, n+1)

	mem[0], mem[1], mem[2] = 0, 1, 1

	if n < 3 {
		return mem[n]
	}

	for i := 3; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2] + mem[i-3]
	}

	return mem[n]
}

func main() {

	fmt.Println(tribonacci(5))
}
