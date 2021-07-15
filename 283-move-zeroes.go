package main

import "fmt"

/**
 * 给定一个数组 nums
 * 编写一个函数将所有 0 移动到数组的末尾
 * 同时保持非零元素的相对顺序。
 */

func main() {
	var nums []int
	nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// 使用双指针
// 一次遍历
func moveZeroes(nums []int) {
	var (
		i, j, length int = 0, 0, len(nums)
	)

	// i指针不断向后遍历, 遇到非零立刻交换
	// j指针指向左侧为0位置,被i指针交换后自增, 向后移动
	for i < length {
		// 元素不为0
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
		i++
	}
	fmt.Println(nums)
}

/**
 * 两次遍历
 * 第一次遍历移动
 * 第一次遍历补全0
 */
func moveZeroes1(nums []int) {
	var (
		i, j, length int = 0, 0, len(nums)
	)

	for i < length {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	for j < length {
		nums[j] = 0
		j++
	}
	fmt.Println(nums)
}
