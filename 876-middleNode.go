// Package main provides ...
package main

import "fmt"

/**
 * 链表的中间节点
 */
func main() {
	var nums []int
	nums = []int{7}
	nums = []int{5, 7}
	//nums = []int{1, 2, 3, 4, 5}
	//nums = []int{1, 2, 3, 4, 5, 6}
	//nums = []int{1, 2, 3, 4, 5, 6, 7}

	list := newList(nums)
	r := middleNode(list)
	fmt.Println("mid node is: ", r.Val)
}

func newList(nums []int) *ListNode {
	var r *ListNode = newNode(0)
	head := r
	for _, v := range nums {
		r.Next = newNode(v)
		r = r.Next
	}
	return head.Next
}

func newNode(n int) *ListNode {
	return &ListNode{
		Val:  n,
		Next: nil,
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var fast, slow *ListNode = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
