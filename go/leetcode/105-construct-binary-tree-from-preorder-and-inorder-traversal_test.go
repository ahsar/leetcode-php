// 中序&后序还原二叉树
package leetcode

import (
	. "algo/bintree"
	"testing"
)

func Test105(*testing.T) {
	var preorder, inorder []int

	// case 1
	preorder = []int{3, 9, 20, 15, 7}
	inorder = []int{9, 3, 15, 20, 7}

	// case 2
	preorder = []int{3, 9, 15, 13, 11, 17, 12, 20, 16, 21}
	inorder = []int{13, 15, 11, 9, 12, 17, 3, 16, 20, 21}

	tree := buildTree2(preorder, inorder)
	LevelOrder(tree)
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	var node *TreeNode
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	root := preorder[0]
	// 前序: 根->左->右
	// 中序: 左->根->右
	var index int
	for ; index < len(inorder); index++ {
		if inorder[index] == root {
			break
		}
	}

	// 用计算去划分区间
	node = &TreeNode{
		Val:   root,
		Left:  buildTree2(preorder[1:index+1], inorder[:index]),
		Right: buildTree2(preorder[index+1:], inorder[index+1:]),
	}
	return node
}
