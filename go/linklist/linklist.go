// Package linklist provides ...
package linklist

import "fmt"

func NewList(nums []int) *ListNode {
	var r *ListNode = newNode(0)
	head := r
	for _, v := range nums {
		r.Next = newNode(v)
		r = r.Next
	}
	return head.Next
}

func PrintLink(head *ListNode) {
	for head != nil {
		//fmt.Printf("%d -> ", head.Val)
		fmt.Printf("%d {%p} -> ", head.Val, head)
		head = head.Next
	}
	fmt.Print("nil\n")
}

func newNode(n int) *ListNode {
	return &ListNode{
		Val:  n,
		Next: nil,
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
