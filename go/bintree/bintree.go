package bintree

import (
	"algo/stack"
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
	if i > l-1 || arr[i] == 0 {
		return nil
	}

	t := NewRoot()
	t.Val = arr[i]
	t.Left = BuildTree(2*i+1, l, arr)
	t.Right = BuildTree(2*i+2, l, arr)
	return t
}

func LevelOrder(root *TreeNode) {
	s := stack.NewStack()
	s.Push(root)

	var (
		tmp *TreeNode
		ok  bool
	)

	for !s.Empty() {
		if tmp, ok = s.Pop().(*TreeNode); !ok {
			fmt.Println("level order err:", ok)
			continue
		}

		if tmp.Left != nil {
			fmt.Println(tmp.Val, "left is", tmp.Left.Val)
			s.Push(tmp.Left)
		}
		if tmp.Right != nil {
			fmt.Println(tmp.Val, "right is", tmp.Right.Val)
			s.Push(tmp.Right)
		}
	}
}
