package leetcode

import (
	"fmt"
	"testing"
)

func Test200(*testing.T) {
	var grid [][]byte
	var board [][]int

	// case 1
	board = [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// case 2
	board = [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	grid = make([][]byte, len(board))
	for i := 0; i < len(board); i++ {
		grid[i] = make([]byte, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			grid[i][j] = byte(board[i][j])
		}
	}

	r := numIslands(grid)
	fmt.Println("res:", r)
}

/**
 * 岛屿的数量(深度优先遍历)
 */
func numIslands(grid [][]byte) int {
	var printLand func(int, int)
	printLand = func(i, j int) {
		if grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		// 向上
		if i > 0 && grid[i-1][j] == 1 {
			printLand(i-1, j)
		}
		// 向下
		if i < len(grid)-1 && grid[i+1][j] == 1 {
			printLand(i+1, j)
		}
		// 向右
		if j < len(grid[i])-1 && grid[i][j+1] == 1 {
			printLand(i, j+1)
		}
		// 向左
		for j > 0 && grid[i][j-1] == 1 {
			printLand(i, j-1)
		}
	}

	var r int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			printLand(i, j)
			r++
		}
	}

	return r
}
