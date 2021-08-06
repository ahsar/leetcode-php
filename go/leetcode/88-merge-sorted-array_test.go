package leetcode

import (
	"fmt"
	"testing"
)

func Test88(*testing.T) {
	var (
		nums1, nums2 []int
		m, n         int
	)

	nums1 = []int{1, 2, 3, 0, 0, 0}
	m = 3
	nums2 = []int{2, 5, 6}
	n = 3

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		x    int = len(nums1) - 1
		i, j int = m - 1, n - 1
	)

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[x] = nums1[i]
			i--
		} else {
			nums1[x] = nums2[j]
			j--
		}
		x--
	}

	for i >= 0 {
		nums1[x] = nums1[i]
		x--
		i--
	}

	for j >= 0 {
		nums1[x] = nums2[j]
		x--
		j--
	}
	fmt.Println(nums1)
}
