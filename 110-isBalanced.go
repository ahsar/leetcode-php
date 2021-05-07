// Package main provides ...
package main

/**
 * 平衡二叉树
 */
func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return abs(height(root.Left), height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if nil == root {
		return 0
	}

	lLen := height(root.Left)
	rLen := height(root.Right)
	if lLen > rLen {
		return lLen + 1
	} else {
		return rLen + 1
	}
}

func abs(i int, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}
