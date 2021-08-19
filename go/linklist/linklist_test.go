package linklist

import (
	"testing"
)

func TestLinkIsert(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4, 5}
	nums = []int{3, 4, 6, 7, 8}
	link := NewList(nums)
	head := link
	link.Next.Next.Next.Next = head

	node := CycleLinkInsert(link, 2)
	PrintCycleLink(node)
}

func CycleLinkInsert(root *ListNode, val int) *ListNode {
	var head *ListNode = root

	node := &ListNode{Val: val}
	// 空链表指向自己
	if root == nil {
		node.Next = node
		return node
	}

	prev := root
	cur := root.Next
	for cur != head {
		if prev.Val <= val && val <= cur.Val {
			break
		}
		prev = cur
		cur = cur.Next
	}
	prev.Next = node
	node.Next = cur

	if val > head.Val {
		return head
	}
	return node
}
