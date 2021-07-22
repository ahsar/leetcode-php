// Package leetcode provides ...
package leetcode

import (
	"testing"

	. "algo/linklist"
)

// 删除链表的倒数第 N 个结点
func Test19(t *testing.T) {
	var (
		n    int
		nums []int
	)

	//case 1
	nums = []int{1, 2, 3, 4, 5}
	n = 5

	//case 2
	nums = []int{1}
	n = 1

	//case 3
	nums = []int{1, 2}
	n = 2

	list := NewList(nums)
	r := removeNthFromEnd(list, n)
	PrintLink(r)
}

/**
 * (双指针)
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode = head, head

	// 快指针先走n+1步
	for i := 0; i <= n; i++ {
		// 如果快指针直接到了终点, 证明n>=链表长度
		if fast == nil {
			return head.Next
		}
		fast = fast.Next
	}

	// 快指针到头,此时满指针在倒数n+1处
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 释放对象
	var next *ListNode = slow.Next.Next
	slow.Next = nil
	slow.Next = next
	return head
}
