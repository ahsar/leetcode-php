// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"

	. "algo/linklist"
)

func middleNode(head *ListNode) *ListNode {
	var fast, slow *ListNode = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

/**
 * 链表的中间节点
 */
func Test876(t *testing.T) {
	var nums []int
	nums = []int{7}
	nums = []int{5, 7}
	//nums = []int{1, 2, 3, 4, 5}
	//nums = []int{1, 2, 3, 4, 5, 6}
	//nums = []int{1, 2, 3, 4, 5, 6, 7}

	list := NewList(nums)
	r := middleNode(list)
	fmt.Println("mid node is: ", r.Val)
}
