// 二叉树锯齿层序遍历
package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test103(*testing.T) {
	var nums []int

	// case 1
	nums = []int{3, 9, 20, 0, 0, 15, 7}

	// case 2
	nums = []int{1}

	// case 3
	nums = []int{1, 2, 3}

	tree := BuildTree(0, len(nums), nums)
	r := zigzagLevelOrder(tree)
	fmt.Println(r)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var (
		que      []*TreeNode
		i        int
		tmpLevel = []int{}
	)

	que = append(que, root)
	for len(que) > 0 {
		level := que
		que = nil
		tmpLevel = nil

		for _, r := range level {
			tmpLevel = append(tmpLevel, r.Val)
			if r.Left != nil {
				que = append(que, r.Left)
			}
			if r.Right != nil {
				que = append(que, r.Right)
			}
		}

		// 奇数从右向左
		if (i % 2) != 0 {
			for i, j := 0, len(tmpLevel)-1; i < j; i, j = i+1, j-1 {
				tmpLevel[i], tmpLevel[j] = tmpLevel[j], tmpLevel[i]
			}
		}
		i++
		res = append(res, tmpLevel)
	}

	return res
}
