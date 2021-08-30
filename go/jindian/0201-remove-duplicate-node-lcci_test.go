// Package jindian provides ...
// 移除重复元素
package jindian

import (
	. "algo/linklist"
	"testing"
)

func Test0201(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 3, 2, 1}
	link := NewList(nums)
	//PrintLink(link)
	removeDuplicateNodes(link)
	PrintLink(link)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	var (
		node *ListNode    = head
		hash map[int]bool = make(map[int]bool)
	)

	hash[head.Val] = true
	for node.Next != nil {
		if hash[node.Next.Val] {
			node.Next = node.Next.Next
			continue
		}

		hash[node.Next.Val] = true
		node = node.Next
	}
	return head
}
