package leetcode

import (
	"fmt"
	"testing"
)

func Test62(*testing.T) {
	var m, n int

	// case 1
	m = 3
	n = 7

	// case 2
	m = 3
	n = 3

	// case 3
	m = 3
	n = 2

	// case 4
	m = 1
	n = 1

	r := uniquePaths(m, n)
	fmt.Println(r)
}

func uniquePaths(m int, n int) int {
	var dp [][]int
	dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
