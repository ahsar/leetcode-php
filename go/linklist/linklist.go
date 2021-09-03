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

/**
 * 打印链表
 */
func PrintLink(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		//fmt.Printf("%d {%p} -> ", head.Val, head)
		head = head.Next
	}
	fmt.Print("nil\n")
}

/**
 * 链表逆序打印
 */
func ReversePrint(head *ListNode) {
	if head == nil {
		return
	}

	ReversePrint(head.Next)
	fmt.Printf("%d {%p} <- ", head.Val, head)
}

/**
 * 打印环形链表
 */
func PrintCycleLink(head *ListNode) {
	var root *ListNode = head

	for root != head.Next {
		fmt.Printf("%d {%p} -> ", head.Val, head)
		head = head.Next
	}
	fmt.Printf("%d {%p} -> %d {%p} ", head.Val, head, head.Next.Val, head.Next)
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
