package jindian

import (
	. "algo/linklist"
	"fmt"
	"testing"
)

func Test0206(*testing.T) {
	var nums []int

	//case1
	nums = []int{1, 2, 3, 4, 3, 2, 1}

	//case2
	nums = []int{1, 2, 3, 3, 2, 1}

	//case3
	nums = []int{1}

	//case4
	nums = []int{1, 1}

	//case5
	nums = []int{1, 3, 2, 1}

	link := NewList(nums)
	PrintLink(link)
	r := isPalindrome(link)
	fmt.Println("res", r)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	tmp = head
	return check(head)
}

var tmp *ListNode

func check(head *ListNode) bool {
	if head == nil {
		return true
	}

	defer func() {
		tmp = tmp.Next
	}()

	r := check(head.Next)
	return r && head.Val == tmp.Val
}
