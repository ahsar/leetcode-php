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

func moveZeroes(nums []int) {
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
