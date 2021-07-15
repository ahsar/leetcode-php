// Package main provides ...
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 双指针题型
package main

import (
	"fmt"
)

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
	nums = []int{-1}
	k = 2

	//case 4
	//nums = []int{1, 2}
	//k = 3

	//case 5
	//nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//k = 3

	fmt.Println("origin: ", nums)
	rotate(nums, k)
}

/**
 * 翻转法
 */
func rotate(nums []int, k int) {
	k = k % len(nums)
	// 整体反转
	reverse(nums)

	reverse(nums[0:k])
	reverse(nums[k:])
	fmt.Println(nums)
}

// 数组反转函数
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

// (环状替换)
// 依次轮转替换, 直到替换次数=数组长度
// todo (轮转次数)
func rotate2(nums []int, k int) {
	length := len(nums)
	k = k % length

	var (
		prev  int     // 两两交换上一次值
		next  int     // 两两交换下一次值
		n     int     // 记录交换次数
		i     int = 0 // 当前索引
		mod   int
		extra int = 0
	)

	// 元素的第一个元素
	prev = nums[i]
	for n < length+extra {
		mod = (i + k) % length
		// 保存被覆盖的后k位
		next = nums[mod]
		// 后k位等于前k位元素
		nums[mod] = prev
		// 此时下一个元素等于上一次循环的元素
		prev = next
		i += k

		// 如果替换回到原始位置, 则向后移位
		if mod == 0 {
			//extra++
			i++
		}
		fmt.Println(i, mod, prev, nums)
		n++
	}
	fmt.Println("res: ", nums)
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
