// Package leetcode provides ...
// 在排序数组中查找元素的第一个和最后一个位置
package leetcode

import (
	"fmt"
	"testing"
)

func Test34(*testing.T) {
	var (
		nums   []int
		target int
	)

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 8

	nums = []int{7, 7, 7, 8, 8, 10}
	target = 7

	nums = []int{7, 7, 7, 7}
	target = 7

	//nums = []int{1}
	//target = 0

	//nums = []int{1}
	//target = 1

	//nums = []int{2, 2}
	//target = 3

	nums = []int{1, 2}
	target = 1

	r := searchRange(nums, target)
	fmt.Println(r)
}

/**
 *  二分查询变种
 */
func searchRange(nums []int, target int) []int {
	var (
		mid    int
		res    []int = make([]int, 2)
		length int   = len(nums)
	)
	res = []int{-1, -1}
	if length < 1 {
		return res
	}

	//todo
	for l, r := 0, length; l <= r; {
		mid = (l + r) / 2
		if target == nums[mid] {
			// 如果相等,则向左侧自旋直到找到最左
			for ; mid >= 0 && target == nums[mid]; mid-- {
			}
			// 如果相等,则向右侧侧自旋直到找到最左
			mid++
			res[0] = mid
			for ; mid < length && target == nums[mid]; mid++ {
			}
			mid--
			res[1] = mid
			break
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
