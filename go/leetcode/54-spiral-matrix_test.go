// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 螺旋打印数组
 */
func TestSpiralOrder(*testing.T) {
	var matrix [][]int

	// case 1
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// case 2
	//matrix = [][]int{
	//{1, 2, 3, 4},
	//{5, 6, 7, 8},
	//{9, 10, 11, 12},
	//}

	// case 3
	//matrix = [][]int{{1}}

	// case 4
	//matrix = [][]int{{1, 2}}

	// case 5
	//matrix = [][]int{
	//{1},
	//{2},
	//}

	// case 6
	//matrix = [][]int{
	//{1},
	//{2},
	//{3},
	//}

	// case 7
	//matrix = [][]int{
	//{1, 2},
	//{3, 4},
	//{5, 6},
	//}

	r := spiralOrder(matrix)
	fmt.Println("res", r)
}

/**
 *  顺时针打印数组
 */
func spiralOrder(matrix [][]int) []int {
	var (
		up    int = 0
		left  int = 0
		right int = len(matrix[0]) - 1
		down  int = len(matrix) - 1

		path  int   = 0
		times int   = len(matrix[0]) * len(matrix)
		res   []int = make([]int, times)
	)

LOOP:
	for times > path {
		// 左->右
		for i := left; i <= right; i++ {
			res[path] = matrix[up][i]
			path++
			if path >= times {
				break LOOP
			}
		}
		// 行结束, 上边界收缩
		up++

		// 上->下
		for j := up; j <= down; j++ {
			res[path] = matrix[j][right]
			path++
			if path >= times {
				break LOOP
			}
		}
		// 上下列结束, 右边界定收缩
		right--

		// 右->左
		for x := right; x >= left; x-- {
			if path >= times {
				break LOOP
			}
			res[path] = matrix[down][x]
			path++
		}
		down--

		// 下->上
		for y := down; y >= up; y-- {
			res[path] = matrix[y][left]
			path++
			if path >= times {
				break LOOP
			}
		}
		left++
	}

	return res
}
