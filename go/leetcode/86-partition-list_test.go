// Package leetcode provides ...
package leetcode

/**
 * 分隔链表
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 *
 * case 1
 * link: 1 -> 4 -> 3 -> 2 -> 5 -> 2
 * x: 3
 * new: 1 -> 2 -> 2 -> 4 -> 3 -> 5
 *
 * case 2
 * link: 1 -> 10 -> 3 -> 4 -> 2 -> 9
 * x: 4
 * new: 1 -> 3 -> 2 -> 10 -> 4 -> 9
 *
 * skill:
 * 使用small + large 两个链表进行分别保存然后链表拼接
 *
 */

import (
	. "algo/linklist"
	"testing"
)

func Test86(*testing.T) {
	var (
		nums []int
		x    int
	)

	// case 1
	nums = []int{1, 4, 3, 2, 5, 2}
	x = 2

	// case 2
	nums = []int{1, 4, 3, 2, 5, 2}
	x = 3

	// case 3
	nums = []int{2, 1}
	x = 2

	// case 4
	nums = []int{-2, 1}
	x = 2

	link := NewList(nums)
	PrintLink(link)
	link = partition(link, x)
	PrintLink(link)
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var (
		small     *ListNode = &ListNode{}
		large     *ListNode = &ListNode{}
		smallHead *ListNode
		largeHead *ListNode
	)
	smallHead = small
	largeHead = large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}

		head = head.Next
	}
	large.Next = nil

	small.Next = largeHead.Next
	return smallHead.Next
}
