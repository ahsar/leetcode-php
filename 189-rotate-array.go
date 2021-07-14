// Package main provides ...
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 双指针题型
package main

import "fmt"

func main() {
	var (
		nums []int
		k    int
	)

	//case 1
	// 1, 2, 3, 4, 5, 6, 7
	// 7, 6, 5, 4, 3, 2, 1
	// 5, 6, 7, 1, 2, 3, 4
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3

	//case 2
	//nums = []int{-1, -100, 3, 99}
	//k = 2

	//case 3
	//nums = []int{-1}
	//k = 2

	//nums = []int{1, 2}
	//k = 3

	fmt.Println("origin: ", nums)
	rotate(nums, k)
}

// 1, 2, 3, 4, 5, 6, 7
/**
 * (环状替换)
 * 依次轮转替换, 直到替换次数=数组长度
 */
func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length

	var (
		tmp int     // 两两交换临时变量
		n   int     // 记录交换次数
		i   int = 0 // 当前索引
	)

	for n < length {
		// 当前索引的后k个位置
		tmp = nums[i+k]

		// 当前值向后移动
		nums[i+k] = nums[n]
	}
}

// (暴力法)
// 时间: O(n) + O(k)
// 空间:O(k)
func rotate1(nums []int, k int) {
	length := len(nums)
	k = k % length
	var (
		r          int = length - 1
		i          int = 0
		extra          = make([]int, k)
		extraIndex int = 0
	)

	// 从后向前遍历数组
	// 小于k, 追加到暂存数组中
	// 大于k, 数组向后搬移
	for r >= 0 {
		if i < k {
			extra[extraIndex] = nums[r]
			extraIndex++
		} else {
			// 后第k项等于当前项
			nums[r+k] = nums[r]
		}
		r--
		i++
	}

	i = 0
	for _, v := range extra {
		nums[k-i-1] = v
		i++
	}
	fmt.Println("res: ", nums)
}
