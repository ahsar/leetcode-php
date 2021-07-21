// Package leetcode provides ...
package leetcode

import (
	"fmt"
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
	n = 2

	//case 2
	//nums = []int{1}
	//n = 1

	//case 3
	//nums = []int{1, 2}
	//n = 1

	list := NewList(nums)
	r := removeNthFromEnd(list, n)
	fmt.Println(r)
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
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
