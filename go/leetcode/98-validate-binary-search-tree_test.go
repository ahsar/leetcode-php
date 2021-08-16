package leetcode

import (
	"fmt"
	"math"
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
	//nums = []int{5, 1, 4, 0, 0, 3, 6}

	tree := BuildTree(0, len(nums), nums)
	LevelOrder(tree)
	r := isValidBST(tree)
	fmt.Println(r)
}

/**
 * 验证是否是一颗二叉搜索树
 *
 * 就很巧妙
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(*TreeNode, int, int) bool
	helper = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}

		// 左侧节点不能大于当前节点
		return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}
