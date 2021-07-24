// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test695(t *testing.T) {
	var grid [][]int

	grid = [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	r := maxAreaOfIsland(grid)
	fmt.Println(grid)
	fmt.Println(r)
}

// 深度优先遍历
// 遍历二维数组每一个节点, 遍历过程中如果遇到2, 则停止遍历
// 一旦遍历过, 则对当前节点进行"上色"
func maxAreaOfIsland(grid [][]int) int {
	var (
		l1 int = len(grid)
		l2 int = len(grid[0])
		r  int
	)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			r = max(r, dfs(grid, i, j))
		}
	}
	return r
}

func dfs(grid [][]int, i, j int) int {
	var (
		l1                    int = len(grid)
		l2                    int = len(grid[0])
		left, right, down, up int
	)

	// 已经被上色或者是"水"
	if grid[i][j] == 2 || grid[i][j] == 0 {
		return 0
	}
	// 执行上色
	grid[i][j] = 2

	if i > 0 {
		left = dfs(grid, i-1, j)
	}
	if i < l1-1 {
		right = dfs(grid, i+1, j)
	}
	if j > 0 {
		down = dfs(grid, i, j-1)
	}
	if j < l2-1 {
		up = dfs(grid, i, j+1)
	}

	return 1 + left + right + down + up
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
