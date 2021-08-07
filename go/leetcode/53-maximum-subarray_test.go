package leetcode

import (
	"fmt"
	"testing"
)

func Test53(*testing.T) {
	var nums []int
	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums = []int{1}
	nums = []int{0}
	nums = []int{-1000}
	r := maxSubArray(nums)
	fmt.Println(r)
}

func maxSubArray(nums []int) int {
	var (
		dp  []int
		l   = len(nums)
		res int
	)
	dp = make([]int, l)

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	dp[0] = nums[0]
	res = nums[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}
