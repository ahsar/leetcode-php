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
	r := rob(nums)
	fmt.Println(r)
}

func rob(nums []int) int {
	l := len(nums)

	var r int

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 0; i < l-3; i = i + 2 {
		// 隔1个, 隔2个
		r += max(nums[i]+nums[i+2], nums[i]+nums[i+3])
		fmt.Println(r, "|", max(nums[i]+nums[i+2], nums[i]+nums[i+3]))
	}
	return r
}
