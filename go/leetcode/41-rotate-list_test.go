// Package leetcode provides ...
package leetcode

import (
	. "algo/linklist"
	"fmt"
	"testing"
)

func TestRotateRight(*testing.T) {
	var (
		nums []int
		k    int
	)

	//case 1
	nums = []int{1, 2, 3, 4, 5}
	k = 1

	// case 2
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8}
	k = 9

	// case 3
	nums = []int{1, 2, 3}
	k = 8

	// case 4
	nums = []int{}
	k = 8

	link := NewList(nums)
	PrintLink(link)
	fmt.Println("k:", k)
	link = rotateRight(link, k)
	PrintLink(link)
}

/**
 * 1. 先获取长度
 * 2. 取模移动
 *
 */
func rotateRight(head *ListNode, k int) *ListNode {
	var (
		newHead *ListNode
		dummy   *ListNode = head
		fast    *ListNode = head
		slow    *ListNode = head
		length  int
		cHead   *ListNode = head
	)

	for cHead != nil {
		length++
		cHead = cHead.Next
	}
	if length == 0 {
		return nil
	}
	k = k % length
	if k == 0 {
		return head
	}

	for i := 0; fast.Next != nil && i < k; fast = fast.Next {
		i++
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	newHead = slow.Next
	slow.Next = nil

	// 链表回环
	fast.Next = dummy
	return newHead
}
