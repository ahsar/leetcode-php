package leetcode

import (
	"fmt"
	"testing"
)

func Test33(*testing.T) {
	var nums []int
	var target int

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 4

	nums = []int{8, 1, 2, 3, 4, 5, 6, 7}
	target = 6
	r := search(nums, target)
	fmt.Println(r)
}

/**
 * 搜索旋转排序数组
 *
 * 二分查找
 */
func search(nums []int, target int) int {
	var mid int
	for i, j := 0, len(nums)-1; i <= j; {
		mid = (i + j) / 2
		if target == nums[mid] {
			return mid
		}
		if target == nums[i] {
			return i
		}
		if target == nums[j] {
			return j
		}

		// 左侧范围有序
		if nums[i] < nums[mid] {
			if target > nums[i] && target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else {
			// 右侧范围有序
			if nums[mid] < target && target < nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return -1
}
