package leetcode

import (
	. "algo/linklist"
	"testing"
)

/**
 * 重排链表
 *
 * 1.
 */

func Test143(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 2, 3, 4}

	// case 2
	//nums = []int{1, 2, 3, 4, 5, 6}

	// case 3
	//nums = []int{1, 2, 3, 4, 5}

	// case 4
	nums = []int{1, 2, 5}

	// case 5
	nums = []int{15}

	link := NewList(nums)
	PrintLink(link)
	reorderList(link)
	PrintLink(link)
}

/**
 * 递归法(最好有迭代实现)
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var (
		node       *ListNode = &ListNode{}
		head2      *ListNode
		helper     func(*ListNode)
		fast, slow *ListNode = head, head // slow 用来记录中间节点位置
	)
	head2 = node

	// 1. 获取链表中间节点
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var flag bool
	//2. 获取逆序, 一旦遍历到中点, 后面无需遍历
	helper = func(head *ListNode) {
		if head == nil {
			return
		}

		helper(head.Next)

		if head == slow || flag {
			flag = true
			return
		}

		node.Next = head
		node = node.Next
	}
	helper(head)
	node.Next = nil
	head2 = head2.Next

	// 3. 拼接
	var next1, next2 *ListNode
	for head != nil && head2 != nil {
		next1 = head.Next
		next2 = head2.Next

		head.Next = head2
		head = next1

		head2.Next = head
		head2 = next2
	}
	head.Next = nil
}
