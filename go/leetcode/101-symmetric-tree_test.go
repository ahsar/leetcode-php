// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"

	. "algo/bintree"
)

func Test101(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 2, 2, 3, 4, 4, 3}

	// case 2
	nums = []int{1, 2, 2, 0, 3, 0, 3}

	// case 2
	//nums = []int{1, 2, 3}

	tree := BuildTree(0, len(nums), nums)
	LevelOrder(tree)

	r := isSymmetric(tree)
	fmt.Println(r)
}

/**
 * 对称二叉树
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var checkSymmetric func(*TreeNode, *TreeNode) bool
	checkSymmetric = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		return checkSymmetric(left.Left, right.Right) && checkSymmetric(left.Right, right.Left)
	}

	return checkSymmetric(root.Left, root.Right)
}
