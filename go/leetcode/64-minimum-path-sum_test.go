// 最小路径和(简单)
package leetcode

import (
	"fmt"
	"testing"
)

func Test64(*testing.T) {
	var grid [][]int

	// case 1
	grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	// case 2
	//grid = [][]int{{1, 2, 3}, {4, 5, 6}}

	r := minPathSum(grid)
	fmt.Println(r)
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j > 0 { // 顶行
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
			if j == 0 && i > 0 { // 左
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
			if j > 0 && i > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
