package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test129(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3}
	nums = []int{4, 9, 0, 5, 1}
	tree := BuildTree(0, len(nums), nums)
	sumNumbers(tree)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	var (
		backTrack func(*TreeNode)
		path      []int
		res       [][]int
		tmp       []int
	)

	push := func(i int) {
		//fmt.Println("push:", i)
		path = append(path, i)
	}

	pop := func() int {
		fmt.Print("pop ")
		defer func() {
			path = path[:len(path)-1]
			fmt.Println("pop:", path)
		}()

		return path[len(path)-1]
	}

	backTrack = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			tmp = make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		push(root.Val)
		if root.Left != nil {
			fmt.Println("push:", root.Val, root.Left)
			backTrack(root.Left)
		}

		if root.Right != nil {
			backTrack(root.Right)
		}

		pop()
	}

	backTrack(root)

	fmt.Println("res:", res)
	return 1
}
