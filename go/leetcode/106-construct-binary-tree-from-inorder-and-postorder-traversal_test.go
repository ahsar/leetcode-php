// 根据中序&后序遍历还原二叉树
package leetcode

import (
	. "algo/bintree"
	"testing"
)

func Test106(*testing.T) {
	var (
		inorder   []int
		postorder []int
	)

	// case 1      左  根 右
	inorder = []int{9, 3, 15, 20, 7}
	postorder = []int{9, 15, 7, 20, 3}

	// case 2      左  根 右
	inorder = []int{21, 11, 23, 9, 17, 3, 15, 20, 16}
	postorder = []int{21, 23, 11, 17, 9, 15, 16, 20, 3}

	tree := buildTree(inorder, postorder)
	LevelOrder(tree)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) < 1 || len(inorder) < 1 {
		return nil
	}

	var node *TreeNode
	// 中序: 左->根->右
	// 中序: 左->右->根
	// 右子树根节点等于后序遍历根节点前一节点元素
	root := postorder[len(postorder)-1]
	// 根据中序遍历, 找到根索引
	var i int = 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root {
			break
		}
	}

	node = &TreeNode{
		Val: root,
		// 利用中序遍历+root确定左长度, 套进后序遍历
		Left: buildTree(inorder[:i], postorder[:i]),
		// 根节点+1到结尾为右子树全部元素,
		Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
	}

	return node
}
