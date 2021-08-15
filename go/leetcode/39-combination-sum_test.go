// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test39(*testing.T) {
	var (
		candidates []int
		target     int
	)

	// case 1
	candidates = []int{2, 8, 3, 6, 7}
	target = 7

	// case 2
	candidates = []int{2, 3, 5}
	target = 8

	// case 3
	candidates = []int{2}
	target = 1

	// case 4
	candidates = []int{1}
	target = 1

	// case 5
	candidates = []int{1}
	target = 2

	// case 6
	candidates = []int{9, 2, 3, 5}
	target = 8

	r := combinationSum(candidates, target)
	fmt.Println(r)
}

func combinationSum(candidates []int, target int) (res [][]int) {
	var path []int
	push := func(i int) {
		path = append(path, i)
	}

	pop := func() int {
		defer func() {
			path = path[:len(path)-1]
		}()

		return path[len(path)-1]
	}

	var (
		total        int
		tmp          []int
		backTracking func(int)
		length       int = len(candidates)
	)

	backTracking = func(startIndex int) {
		if total > target {
			return
		}

		if total == target {
			tmp = make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := startIndex; i < length; i++ {
			if candidates[i] > target {
				continue
			}
			total += candidates[i]
			push(candidates[i])
			backTracking(i)
			pop()
			total -= candidates[i]
		}
	}

	backTracking(0)
	return res
}
