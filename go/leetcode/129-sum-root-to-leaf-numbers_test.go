package leetcode

import (
	. "algo/bintree"
	"fmt"
	"strconv"
	"testing"
)

func Test129(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3}
	//nums = []int{4, 9, 0, 5, 1}
	tree := BuildTree(0, len(nums), nums)
	r := sumNumbers(tree)
	fmt.Println("res:", r)
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
		path      string
		res       int
		num       int
	)

	backTrack = func(root *TreeNode) {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			num, _ = strconv.Atoi(path)
			res += num
			path = path[:len(path)-1]
			return
		}

		if root.Left != nil {
			backTrack(root.Left)
		}

		if root.Right != nil {
			backTrack(root.Right)
		}
		path = path[:len(path)-1]
	}

	backTrack(root)
	return res
}
