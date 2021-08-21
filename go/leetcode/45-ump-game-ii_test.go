// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test45(*testing.T) {
	var nums []int

	// case 1
	nums = []int{2, 3, 1, 1, 4}

	// case 2
	//nums = []int{2, 3, 0, 1, 4}

	// case 3
	//nums = []int{1, 1, 1, 1, 1}

	// case 4
	//nums = []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}

	// case 5
	//nums = []int{1, 2}

	r := jump(nums)
	fmt.Println(r)
}

// 求当前区域的路径最大值, 贪心算法
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var (
		r, cover int = 1, nums[0]
		max      int
		i, j     int
	)

	for i = 0; i <= cover; {
		// 抵达终点(判断抵达终点的要素为当前索引值+可跳步数是否大于数组长度)
		if cover >= len(nums)-1 {
			break
		}

		// 在cover范围内找到路径最大值(贪心)
		max = 0
		for j = i + 1; j <= cover; j++ {
			//fmt.Println(nums[j]+j, nums[j], j)
			if nums[j]+j > max {
				max = nums[j] + j
				i = j
			}
		}

		// 如果找到最大跳跃点,i 会使用区间最大点
		if max > cover {
			cover = nums[i] + i
			r++
		} else {
			fmt.Println("max", max, i)
			i++
		}
	}

	return r
}
