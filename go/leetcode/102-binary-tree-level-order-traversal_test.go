package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test102(*testing.T) {
	var nums []int
	nums = []int{3, 9, 20, 0, 0, 15, 7}
	nums = []int{1}
	tree := BuildTree(0, len(nums), nums)
	r := levelOrder(tree)
	fmt.Println(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var (
		res [][]int
		que []*TreeNode
		tmp []int
	)
	if root == nil {
		return res
	}

	que = append(que, root)
	for len(que) > 0 {
		level := que
		que = nil
		tmp = nil
		for _, node := range level {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}
