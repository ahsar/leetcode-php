package leetcode

import (
	. "algo/linklist"
	"testing"
)

func Test21(*testing.T) {
	var (
		l1, l2     *ListNode
		arr1, arr2 []int
	)

	// case 1
	arr1 = []int{1, 2, 4}
	arr2 = []int{1, 3, 4}

	// case 2
	//arr1 = []int{}
	//arr2 = []int{}

	//// case 3
	//arr1 = []int{}
	//arr2 = []int{0}

	// case 4
	arr1 = []int{1, 3, 5}
	arr2 = []int{2, 4, 4}

	// case 5
	//arr1 = []int{1, 3, 5}
	//arr2 = []int{2, 4, 4, 7, 8, 10}

	// case 6
	//arr1 = []int{1}
	//arr2 = []int{2, 4, 4, 7, 8, 10}

	l1 = NewList(arr1)
	l2 = NewList(arr2)
	r := mergeTwoLists(l1, l2)
	PrintLink(r)
}

/**
 * 合并两个有序链表
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		prev, head *ListNode = &ListNode{}, &ListNode{}
	)
	prev = head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return prev.Next
}
