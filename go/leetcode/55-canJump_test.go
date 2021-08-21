// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test55(*testing.T) {
	var nums []int

	// case 1
	nums = []int{2, 3, 1, 1, 4}

	// case 2
	nums = []int{3, 2, 1, 0, 4}

	// case 3
	nums = []int{4, 2, 1, 0, 4}

	r := canJump(nums)
	fmt.Println(r)
}

/**
 * 跳跃游戏(贪心算法)
 */
func canJump(nums []int) bool {
	var cover int = nums[0]

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 0; i <= cover; i++ {
		fmt.Println("nums", nums[i]+i, nums[i], i)
		cover = max(nums[i]+i, cover)
		if cover >= len(nums)-1 { // cover = 当前位置+当前位置的可跳步数, 到达最后一个即可
			return true
		}
	}
	return false
}
