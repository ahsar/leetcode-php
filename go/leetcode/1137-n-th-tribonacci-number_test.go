package leetcode

import (
	"fmt"
	"testing"
)

func Test1137(*testing.T) {
	var n int
	n = 4
	r := tribonacci(n)
	fmt.Println(r)
}

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	var dp []int = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
