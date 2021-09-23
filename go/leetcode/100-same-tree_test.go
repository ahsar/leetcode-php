// 判断给定两棵二叉树完全相等
package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test100(*testing.T) {
	var (
		nums1, nums2 []int
	)

	// case 1
	nums1 = []int{1, 2, 3, 4, 5}
	nums2 = []int{1, 2, 3, 4, 5}

	// case 2
	nums1 = []int{1, 2, 3, 4, 5}
	nums2 = []int{1, 2, 3, 5}

	tree1 := BuildTree(0, len(nums1), nums1)
	tree2 := BuildTree(0, len(nums2), nums2)
	r := isSameTree(tree1, tree2)
	fmt.Println(r)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//var r bool
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
}
