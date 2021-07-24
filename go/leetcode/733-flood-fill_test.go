package leetcode

import (
	"fmt"
	"testing"
)

/*
 *有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
 *
 *给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。
 *
 *为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。
 */

func Test718(t *testing.T) {
	var (
		image            [][]int
		sr, sc, newColor int
	)

	//case 1
	// 1 1 1
	// 1 1 0
	// 1 0 1
	// [[1,1,1],[1,1,0],[1,0,1]]
	image = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr = 1
	sc = 1
	newColor = 2

	// case 2
	//image = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	//sr = 0
	//sc = 0
	//newColor = 2

	// case 3
	// 0 0 0
	// 0 1 1
	//image = [][]int{{0, 0, 0}, {0, 1, 1}}
	//sr = 1
	//sc = 1
	//newColor = 1

	r := floodFill(image, sr, sc, newColor)
	fmt.Println(r)
}

// 图像渲染
// 递归实现
//todo 迭代法
// 1. 上色
// 2. 向四个方向移动并调用(移动过程避免触达边界以外, 如果新的位置像素和原像素相同给他爱)
// 3. 如果到达边界或已经上色, 则return
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		l1 int = len(image)
		l2 int = len(image[0])
	)

	origin := image[sr][sc]
	// 已经上色了
	if origin == newColor {
		return image
	}
	image[sr][sc] = newColor

	// 向上一维数组索引 -1
	if sr-1 >= 0 && image[sr-1][sc] == origin {
		fmt.Println("A")
		floodFill(image, sr-1, sc, newColor)
	}
	// 向下一维数组索引 +1
	if sr+1 < l1 && image[sr+1][sc] == origin {
		fmt.Println("b")
		floodFill(image, sr+1, sc, newColor)
	}
	// 向👈一维数组索引 -1
	if sc-1 >= 0 && image[sr][sc-1] == origin {
		floodFill(image, sr, sc-1, newColor)
	}
	// 向👉一维数组索引 +1
	if sc+1 < l2 && image[sr][sc+1] == origin {
		floodFill(image, sr, sc+1, newColor)
	}
	return image
}
