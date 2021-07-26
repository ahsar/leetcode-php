package bintree

import "fmt"

type node interface {
	BuildTree()
	LevelOrder()
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (node *Node) NewRoot() *Node {
	return &Node{}
}

func (node *Node) BuildTree(i, l int, arr []int) *Node {
	if i > l-1 || arr[i] == 0 {
		return nil
	}

	t := node.NewRoot()
	t.Val = arr[i]
	t.Left = node.BuildTree(2*i+1, l, arr)
	t.Right = node.BuildTree(2*i+2, l, arr)
	return t
}

func (node *Node) LevelOrder(root *Node) {
	que := []*Node{root}

	for len(que) > 0 {
		tmp := que
		que = nil

		for _, node := range tmp {
			if node.Left != nil {
				que = append(que, node.Left)
				fmt.Println(node.Val, "left is", node.Left.Val)
			}
			if node.Right != nil {
				que = append(que, node.Right)
				fmt.Println(node.Val, "right is", node.Right.Val)
			}
		}
	}
}
