// 最长递增子序列
package leetcode

import (
	"fmt"
	"testing"
)

func Test300(*testing.T) {
	var nums []int

	// case 1
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
	// res: 2 5 7 101

	// case 2
	//nums = []int{0, 1, 0, 3, 2, 3}
	// res: 0 1 2 3

	// case 3
	//nums = []int{7, 7, 7, 7, 7, 7, 7}

	// case 4
	nums = []int{1, 7, 2, 5, 6, 4, 3}

	r := lengthOfLIS(nums)
	fmt.Println("res", r)
}

/**
 * 1. 暴力法动态规划, 双重循环遍历
 * 2. 针对每一层循环进行一次dp
 *
 */
func lengthOfLIS(nums []int) int {
	var (
		res int   = 1
		dp  []int = make([]int, len(nums))
	)

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				// 针对外层循环的当前项dp[i]
				// 等于从当前项开始向前遍历dp数组
				// 如果值小于当前项, 则用max补齐
				// 遍历完成一遍后dp[i]等于前项里边最大值+1
				dp[i] = max(dp[i], dp[j]+1)
				fmt.Println(nums[i], dp[i])
			}
		}
		res = max(res, dp[i])
	}

	return res
}
