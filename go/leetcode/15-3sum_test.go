package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func Test15(*testing.T) {
	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	nums = []int{}
	nums = []int{0, 0, 0, 0}
	nums = []int{-1, 0, 1, 2, -1, -4}
	nums = []int{-1, 1, -1, 1}
	nums = []int{-2, 0, 1, 1, 2}
	nums = []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	r := threeSum(nums)
	fmt.Println(r)
}

// 选定第一个元素+双指针
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var sum int
	for r := range nums {
		// 有序数组, 不可能后边的更小
		if nums[r] > 0 {
			return res
		}
		// 过滤掉外层
		if r > 0 && nums[r] == nums[r-1] {
			continue
		}

		for i, j := r+1, len(nums)-1; i < j; {
			sum = nums[r] + nums[i] + nums[j]
			if sum == 0 {
				res = append(res, []int{nums[r], nums[i], nums[j]})
				// 如果当前项和后一项相同,向后迭代
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				// 向后遍历
				i++
				j--
			} else if sum > 0 {
				j--
			} else {
				i++
			}
		}
	}
	return res
}
