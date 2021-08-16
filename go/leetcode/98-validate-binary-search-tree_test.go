package leetcode

import (
	"fmt"
	"testing"

	. "algo/bintree"
)

func Test98(*testing.T) {
	var nums []int

	// case 1
	nums = []int{2, 1, 3}

	// case 2
	//nums = []int{5, 1, 4, 0, 0, 3, 6}

	// case 3
	//nums = []int{2, 2, 2}

	// case 4
	nums = []int{5, 1, 4, 0, 0, 3, 6}

	tree := BuildTree(0, len(nums), nums)
	LevelOrder(tree)
	r := isValidBST(tree)
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var r bool = true
	if root.Left != nil && root.Left.Val >= root.Val {
		r = false
	}

	if root.Right != nil && root.Right.Val <= root.Val {
		r = false
	}

	return isValidBST(root.Left) && r && isValidBST(root.Right)
}
