// Package jindian provides ...
package jindian

import (
	. "algo/linklist"
	"fmt"
	"testing"
)

func Test0208(*testing.T) {
	var nums []int
	var link *ListNode

	// case 1
	nums = []int{1, 2, 3, 4, 5}
	link = NewList(nums)
	link.Next.Next.Next.Next = link.Next.Next.Next

	// case 2
	//nums = []int{1}
	//link = NewList(nums)

	// case 3
	//nums = []int{1, 2}
	//link = NewList(nums)
	//link.Next.Next = link

	r := detectCycle(link)
	fmt.Println(r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var fast, slow *ListNode = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}

	return nil
}
