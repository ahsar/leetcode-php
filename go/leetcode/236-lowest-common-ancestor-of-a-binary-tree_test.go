package leetcode

import (
	. "algo/bintree"
	"fmt"
	"testing"
)

func Test236(*testing.T) {
	var nums []int

	// case 1
	nums = []int{3, 5, 1, 6, 2, 0, 8, 0, 0, 7, 4}

	tree := BuildTree(0, len(nums), nums)
	p := RandFind(tree, 5)
	q := RandFind(tree, 1)
	//fmt.Printf("%d{%p} %d{%p}\n", p.Val, p, q.Val, q)
	r := lowestCommonAncestor(tree, p, q)
	fmt.Println(r)
}

/**
 * 最小公共祖先
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode

	// 判断左右子树包含p||q || 当前节点就是q,p
	var dfs func(*TreeNode, *TreeNode, *TreeNode) bool
	dfs = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}

		lson := dfs(root.Left, p, q)
		rson := dfs(root.Right, p, q)

		// 判断(节点左||右子树包含p||q节点) || (当前节点为左||右) && (左||右子节点包含p||q)
		if (lson && rson) || (root.Val == p.Val || root.Val == q.Val) && (lson || rson) {
			ans = root
		}
		return lson || rson || (root.Val == p.Val || root.Val == q.Val)
	}

	dfs(root, p, q)
	return ans
}
