package leetcode

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestThreeSumClosest(*testing.T) {
	var (
		nums   []int
		target int
	)
	nums = []int{-1, 2, 1, -4}
	target = 1

	nums = []int{-21, 8, -16, 12, 11, -25, 23, 15, 0, 4}
	target = 3

	r := threeSumClosest(nums, target)
	fmt.Println("res:", r)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var total, dist int = 0, math.MaxInt32
	var abs = func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	var absMin = func(i, j int) int {
		if abs(i) < abs(j) {
			return i
		}
		return j
	}

	for r := range nums {
		for i, j := r+1, len(nums)-1; i < j; {
			total = nums[r] + nums[i] + nums[j]
			if total == target {
				return target
			}

			dist = absMin(dist, total-target)
			if total > target {
				j--
			} else {
				i++
			}
		}
	}

	return dist + target
}
