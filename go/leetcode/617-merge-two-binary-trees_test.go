// Package leetcde provides ...
package leetcode

import (
	"fmt"
	"testing"

	. "algo/bintree"
)

func Test617(t *testing.T) {
	var (
		root1, root2 *TreeNode
		tree1, tree2 []int
	)

	//case 1
	tree1 = []int{1, 3, 2, 5}
	tree2 = []int{0: 2, 1: 1, 2: 3, 4: 4, 6: 7}

	root1 = BuildTree(0, len(tree1), tree1)
	LevelOrder(root1)
	fmt.Println("---")
	root2 = BuildTree(0, len(tree2), tree2)
	LevelOrder(root2)
	//mergeTrees(root1, root2)
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return root2
}
