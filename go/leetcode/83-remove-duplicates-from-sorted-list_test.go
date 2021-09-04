// Package leetcode provides ...
package leetcode

// 83-删除排序链表的重复元素

// case1:
// 1 -> 1 -> 3
// 1 -> 3

import (
	. "algo/linklist"
	"testing"
)

func Test83(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 1, 2}

	// case 2
	nums = []int{1, 1, 2, 3, 3}

	// case 3
	nums = []int{1}

	// case 4
	nums = []int{}

	// case 5
	nums = []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4}

	link := NewList(nums)
	PrintLink(link)
	link = deleteDuplicates1(link)
	PrintLink(link)
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		dummy *ListNode = head
		link  *ListNode
		cur   int
	)
	link = dummy
	for head != nil {
		cur = head.Val
		for head != nil && head.Val == cur {
			head = head.Next
		}

		dummy.Next = head
		dummy = dummy.Next
	}
	return link
}
