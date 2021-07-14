// Package main provides ...
package main

import "fmt"

/**
 * 给定一个排序数组和一个目标值，在数组中找到目标值，
 * 并返回其索引。如果目标值不存在于数组中，
 * 返回它将会被按顺序插入的位置。
 */

func main() {
	var (
		nums   []int
		target int
	)

	//case 1
	nums = []int{1, 3, 5, 6}
	//nums = []int{1, 3}
	target = 5

	//target = 2 //case 2
	//target = 7 //case 3
	//target = 0 //case 4
	//target = 1 //case 5

	r := searchInsert(nums, target)
	fmt.Println(r)
}

func searchInsert(nums []int, target int) int {
	var (
		l, r, mid int = 0, len(nums), 0
	)

	for l <= r {
		mid = (l + r) / 2
		if mid >= len(nums) {
			break
		}

		/*
		 *中间值大于目标值
		 *如果中间值的前一项小于targe, 则停止
		 *否则, 右边界继续自减
		 */
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				r = mid - 1
			}
		} else {
			l = mid + 1
		}
	}

	return mid
}
