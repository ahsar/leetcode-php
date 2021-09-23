// 二叉树中序遍历
package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test94(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 4, 2, 3}

	// case 2
	//nums = []int{1, 2}

	// case 3
	//nums = []int{1, 0, 2}

	tree := BuildTree(0, len(nums), nums)
	LevelOrder(tree)
	r := inorderTraversal(tree)
	fmt.Println("res", r)
}

/**
 * 中序遍历(左根右)
 */
func inorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nil
	}

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		nums = append(nums, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)

	return nums
}
