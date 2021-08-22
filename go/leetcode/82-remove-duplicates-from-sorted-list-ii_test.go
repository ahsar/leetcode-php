// Package leetcode provides ...
package leetcode

// 删除链表连续重复元素

import (
	. "algo/linklist"
	"testing"
)

func Test82(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 2, 3, 3, 4, 4, 5}

	// case 2
	nums = []int{1, 1, 1, 3, 4, 4, 5}

	link := NewList(nums)
	PrintLink(link)
	r := deleteDuplicates(link)
	PrintLink(r)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var dummy *ListNode = &ListNode{Next: head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		// 通过中间变量, 再删除过程中如果碰到相同元素继续后移动
		x := cur.Next.Val
		if cur.Next.Val == cur.Next.Next.Val {
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
