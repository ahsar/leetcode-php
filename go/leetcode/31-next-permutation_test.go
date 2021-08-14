// Package leetcode provides ...
// 下一个排列(寻找下一个比当前大的第一个排列)
package leetcode

import (
	"fmt"
	"testing"
)

func Test31(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3}
	//nums = []int{6, 5, 3, 1}
	//nums = []int{3, 2, 1}
	//nums = []int{1, 1, 5}
	//nums = []int{2, 1}
	//nums = []int{1, 3, 3, 4}
	//nums = []int{1, 3, 3, 1}

	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	length := len(nums)

	var i int
	// 定位交换点
	for i = length - 1; i > 0; i-- {
		// 前一项小于当前项
		if nums[i-1] < nums[i] {
			break
		}
	}

	// 如果全是降序数组, 直接交换, 没有转折点
	if i > 0 {
		// 此时找到最小的大于nums[i-1]的元素
		for l := length - 1; l > 0; l-- {
			if nums[l] > nums[i-1] {
				nums[l], nums[i-1] = nums[i-1], nums[l]
				break
			}
		}
	}

	// 从转折点到最后开始交换
	for j := length - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
