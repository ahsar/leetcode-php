package leetcode

import (
	"fmt"
	"testing"
)

func Test215(*testing.T) {
	var nums []int
	nums = []int{5, 1, 1, 2, 0, 0}
	nums = []int{10, 3, 5, 2, 5, 1}
	nums = []int{6, 1, 3, 9, 8}
	nums = []int{3, 2, 1, 5, 6, 4}

	var k int
	k = 2
	r := findKthLargest(nums, k)
	fmt.Println(nums, r)
}

func findKthLargest(nums []int, k int) int {
	return findK(nums, 0, len(nums)-1, k)
}

func findK(nums []int, l, r, k int) int {
	partion := func(nums []int, l, r int) int {
		privot := nums[r]

		// i 是已处理区间右边界
		// 一旦当前数值小于pivot,就把当前值放到i的右侧
		var i int = l
		for j := l; j < r; j++ {
			// 交换
			if nums[j] > privot {
				if i != j {
					nums[i], nums[j] = nums[j], nums[i]
				}
				i++
			}
		}

		nums[i], nums[r] = nums[r], nums[i]
		return i
	}

	p := partion(nums, l, r)
	if p == k-1 {
		return nums[p]
	}

	if k-1 < p {
		return findK(nums, l, p-1, k)
	} else {
		return findK(nums, p+1, r, k)
	}
}
