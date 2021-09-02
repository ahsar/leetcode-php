// Package leetcode provides ...
package leetcode

import (
	. "algo/linklist"
	"testing"
)

func Test24(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 2, 3, 4}

	// case 2
	nums = []int{1, 2, 3, 4, 5, 6}

	// case 3
	nums = []int{}

	// case 4
	nums = []int{5}

	// case 5
	nums = []int{1, 2}

	// case 6
	nums = []int{1, 2, 3}

	// case 7
	nums = []int{1, 2, 3, 4, 5}

	link := NewList(nums)
	link = swapPairs(link)
	PrintLink(link)
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1 2 3 4
	// 2 1 4 3
	one := head
	two := head.Next
	three := two.Next

	two.Next = one
	one.Next = swapPairs1(three)
	return two
}

/**
 * 两两交换链表节点
 */
// 1 2 3 4
// 2 1 4 3
func swapPairs(head *ListNode) *ListNode {
	var (
		dummy *ListNode = &ListNode{Next: head}
		cur   *ListNode
		node1 *ListNode
		node2 *ListNode
		node3 *ListNode
	)

	cur = dummy
	for head != nil && head.Next != nil {
		node1 = head       // 1
		node2 = head.Next  // 2
		node3 = node2.Next // 3

		// dummy指向node2
		dummy.Next = node2 // d->2
		node2.Next = node1 // 2->1
		node1.Next = node3 // 1->3
		head = node1.Next  // head: 3
		dummy = node1      // dummy:1->3->4, 在3位head情况下进行向后交替
	}

	return cur.Next
}
