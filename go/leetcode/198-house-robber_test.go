package leetcode

import (
	"fmt"
	"testing"
)

func Test198(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 1}
	nums = []int{2, 7, 9, 3, 1}
	nums = []int{101, 7, 9, 3, 1, 2, 4, 99}
	nums = []int{4}
	r := rob(nums)
	fmt.Println(r)
}

func rob(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return nums[0]
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if l <= 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, l)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}
