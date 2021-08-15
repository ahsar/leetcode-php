// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test11(*testing.T) {
	var height []int

	// case 1
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	// case 2
	height = []int{1, 1}

	// case 3
	height = []int{1, 2, 1}

	// case 4
	height = []int{2, 3, 4, 5, 18, 17, 6}

	r := maxArea(height)
	fmt.Println("res:", r)
}

func maxArea(height []int) int {
	var r int

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	//
	for i, j := 0, len(height)-1; i < j; {
		r = max(min(height[j], height[i])*(j-i), r)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return r
}
