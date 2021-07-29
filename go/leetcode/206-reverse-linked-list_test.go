package leetcode

import (
	"testing"

	. "algo/linklist"
)

func Test206(*testing.T) {
	var (
		list *ListNode
		link []int
	)

	//case 1
	link = []int{1, 2, 3, 4, 5}

	//case 2
	//link = []int{1, 2}

	list = NewList(link)
	r := reverseList(list)
	PrintLink(r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var (
		prev, next *ListNode = nil, &ListNode{}
	)

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
