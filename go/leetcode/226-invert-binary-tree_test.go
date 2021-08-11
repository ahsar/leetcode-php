package leetcode

import (
	. "algo/bintree"
	"testing"
)

func Test226(*testing.T) {
	var nums []int
	nums = []int{4, 2, 7, 1, 3, 6, 9}

	tree := BuildTree(0, len(nums), nums)
	r := invertTree(tree)
	LevelOrder(r)
}

// 反转二叉树(没啥可说的)
func invertTree(root *TreeNode) *TreeNode {
	swap(root)
	return root
}

func swap(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	swap(root.Left)
	swap(root.Right)
}
