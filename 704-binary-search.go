// Package main provides ...
package main

import "fmt"

func main() {
	var (
		nums   = []int{-1, 0, 3, 5, 9, 12}
		target int
	)
	nums = []int{-1}

	target = 9 // case 1
	//target = 2 // case 2
	target = -1 // case 2

	r := search(nums, target)
	fmt.Println(r)
}

func search(nums []int, target int) int {
	var (
		l, r, mid int = 0, 0, 0
	)

	r = len(nums) - 1
	fmt.Println(l, r, mid)

	// 左边界超过右边界, 此时需要停止
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		if nums[mid] < target {
			l = mid + 1
			continue
		}
	}

	return -1
}
