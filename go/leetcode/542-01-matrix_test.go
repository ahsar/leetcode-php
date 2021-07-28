package leetcode

import (
	"fmt"
	"testing"
)

func Test542(t *testing.T) {
	var mat [][]int
	mat = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	mat = [][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}
	// [0,1,0,1,1]
	// [1,1,0,0,1]
	// [0,0,0,1,0]
	// [1,0,1,1,1]
	// [1,0,0,0,1]
	r := updateMatrix(mat)
	fmt.Println(r)
}

//dp 方式解决
// 1. 初始化数组(将非1改为极大值)
// 2. 动态规划(mat[i][j] == 1+min(mat[i][j-1], mat[i][j+1], mat[i-1][j], mat[i+1][j]))
// 	2.1 考虑方向, 从边界开始, 对角线扩散, 左边界为动态规划起始地址
// 3. 左上角到右下角
func updateMatrix(mat [][]int) [][]int {
	var (
		r int = len(mat)
		c int = len(mat[0])
	)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] = 10000
		}
	}

	// 左上角开始遍历(从左到右, 从上到下)
	// 0,0 0,1 0,2 || 如果i,j =1,1开始遍历, 则左顶点无法更新
	// 1,0 1,1 1,2
	// 2,0 2,1 2,2
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i > 0 {
				mat[i][j] = min(mat[i-1][j]+1, mat[i][j])
			}
			if j > 0 {
				mat[i][j] = min(mat[i][j-1]+1, mat[i][j])
			}
		}
	}

	// 右下角开始遍历
	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if i < r-1 {
				mat[i][j] = min(mat[i+1][j]+1, mat[i][j])
			}
			if j < c-1 {
				mat[i][j] = min(mat[i][j+1]+1, mat[i][j])
			}
		}
	}

	return mat
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
