// Package main provides ...
/**
 *  二叉树的深度
 */

package main

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
func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	lLen := maxDepth(root.Left)
	rLen := maxDepth(root.Right)

	if lLen > rLen {
		return lLen + 1
	} else {
		return rLen + 1
	}
}
