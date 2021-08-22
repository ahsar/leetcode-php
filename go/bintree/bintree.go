package bintree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewRoot() *TreeNode {
	return &TreeNode{}
}

//https://juejin.cn/post/6844904145346887694
func BuildTree(i, l int, arr []int) *TreeNode {
	if i > l-1 || arr[i] < 1 {
		return nil
	}

	t := NewRoot()
	t.Val = arr[i]
	t.Left = BuildTree(2*i+1, l, arr)
	t.Right = BuildTree(2*i+2, l, arr)
	return t
}

func RandFind(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	var r *TreeNode
	if root.Left != nil {
		v := RandFind(root.Left, val)
		if v != nil {
			r = v
		}
	}
	if root.Right != nil {
		v := RandFind(root.Right, val)
		if v != nil {
			r = v
		}
	}
	return r
}

func LevelOrder(root *TreeNode) {
	que := []*TreeNode{root}

	for len(que) > 0 {
		tmp := que
		que = nil

		for _, node := range tmp {
			if node.Left != nil {
				que = append(que, node.Left)
				fmt.Printf("%d {%p} left is %d {%p} \n", node.Val, node, node.Left.Val, node.Left)
				//fmt.Println(node.Val, "left is", node.Left.Val)
			}
			if node.Right != nil {
				que = append(que, node.Right)
				//fmt.Println(node.Val, "right is", node.Right.Val)
				fmt.Printf("%d {%p} right is %d {%p} \n", node.Val, node, node.Right.Val, node.Right)
			}
		}
	}
}
