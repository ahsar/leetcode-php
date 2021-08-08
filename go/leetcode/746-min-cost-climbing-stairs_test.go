package leetcode

import (
	"fmt"
	"testing"
)

func Test746(*testing.T) {
	var cost []int
	//cost = []int{10, 15, 20}
	cost = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	r := minCostClimbingStairs(cost)
	fmt.Println(r)
}

func minCostClimbingStairs(cost []int) int {
	var (
		l        = len(cost)
		dp []int = make([]int, l)
	)

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < l; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}

	// 取倒数第一步,倒数第二步的最小值
	return min(dp[l-2], dp[l-1])
}
