// Package main provides ...
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 链表有环
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			slow = head
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

func NewNode(val int) *ListNode {
	return &ListNode{val, nil}
}

func main() {
	head := NewNode(3)
	cur := head
	cur.Next = NewNode(2)
	cur = cur.Next
	cur.Next = NewNode(0)

	cur = cur.Next
	cur.Next = NewNode(-4)

	r := detectCycle(head)
	fmt.Println(r)
}
