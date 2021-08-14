package leetcode

import (
	. "algo/linklist"
	"fmt"
	"testing"
)

func Test160(*testing.T) {
	var nums1, nums2 []int

	// case 1
	nums1 = []int{5, 0, 1, 8, 4, 5}
	nums2 = []int{4, 1}

	// case 2
	nums1 = []int{2, 6, 4}
	nums2 = []int{1, 5}

	// case 3
	nums1 = []int{1}
	nums2 = []int{}

	link1 := NewList(nums1)
	link2 := NewList(nums2)
	//setCross(link1, link2, 6)
	PrintLink(link1)
	PrintLink(link2)

	r := getIntersectionNode(link1, link2)
	//fmt.Printf("%d{%p}\n", r.Val, r)
	fmt.Println("cross at", r)
}

/**
 *  把link2合并到link1的指定节点上
 *
 */
func setCross(link1, link2 *ListNode, val int) {
	for link2.Next != nil {
		link2 = link2.Next
	}

	for link1 != nil {
		if link1.Val == val {
			break
		}
		link1 = link1.Next
	}

	link2.Next = link1
}

/**
 * 相交链表
 *
 * 1. 遍历获取A, B 两链表长度
 * 2. 记录更长的一方链表与另一个链表长度差值
 * 3. 指针同时归位&&长的一方先走
 * 4. 短的链表也开始前向运动
 * 5. 一旦两链表节点相等,则终止并返回
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		link1, link2 *ListNode = headA, headB
		l1, l2       int
	)

	// 获取长度
	for headA != nil || headB != nil {
		if headA != nil {
			l1++
			headA = headA.Next
		}
		if headB != nil {
			l2++
			headB = headB.Next
		}
	}

	// 重新归位
	headA = link1
	headB = link2

	extra := l1 - l2
	if extra >= 0 {
		for i := 0; i < extra; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < -extra; i++ {
			headB = headB.Next
		}
	}

	for headA != nil || headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
