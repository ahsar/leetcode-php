// Package leetcode provides ...
// 爬楼梯动态规话
package leetcode

import (
	"fmt"
	"testing"
)

func Test70(*testing.T) {
	var n int

	// case 1
	n = 1
	// case 2
	n = 2
	// case 3
	//n = 10
	r := climbStairs(n)
	fmt.Println(r)
}

// 1 -> 1
// 2 -> 2 (2, 1)
// 3 -> f(2) + f(1)
// 4 -> f(3) + f(2)
/**
 *  动态规划自底向上构建
 *   每次可以走 1 或者 2 步
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var dp []int
	dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
