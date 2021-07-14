// Package main provides ...
// 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
package main

import "fmt"

func main() {
	var nums []int
	nums = []int{-4, -1, 0, 3, 10}
	//nums = []int{-7, -3, 2, 3, 11}
	r := sortedSquares(nums)
	fmt.Println(r)
}

// 使用双指针解决
func sortedSquares(nums []int) []int {
	length := len(nums)

	var (
		l, r, i int = 0, length - 1, length - 1
		res     []int
	)
	res = make([]int, length)

	for l <= r {
		if nums[r]*nums[r] > nums[l]*nums[l] {
			res[i] = nums[r] * nums[r]
			r--
		} else {
			res[i] = nums[l] * nums[l]
			l++
		}
		i--
	}
	return res
}
