// Package main provides ...
package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		// 相遇
		if fast == slow {
			slow = head
			for {
				if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}

	return nil
}
