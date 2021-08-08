package leetcode

import (
	"fmt"
	"testing"
)

func Test136(*testing.T) {
	var nums []int
	nums = []int{2, 2, 1}
	nums = []int{4, 2, 1, 2, 1}
	singleNumber(nums)
}

func Benchmark136(*testing.B) {
	var nums []int
	nums = []int{4, 2, 1, 2, 1}
	r := singleNumber(nums)
	fmt.Println(r)
}

func singleNumber(nums []int) int {
	var single int

	for _, r := range nums {
		single ^= r
	}

	return single
}
