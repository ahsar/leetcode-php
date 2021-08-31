package leetcode

import (
	"fmt"
	"testing"
)

func Test232(*testing.T) {
	var nums []int

	que := Constructor232()
	nums = []int{1, 2, 3, 4, 5}

	var t int
	for _, v := range nums {
		que.Push(v)
		fmt.Println("push:", v)
		t = que.Pop()
		fmt.Println("pop:", t)
	}
	fmt.Println("peek:", que.Peek())
}

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	var num int
	// 栈2不为空, 不断把栈2元素放入栈1
	for len(this.stack2) > 0 {
		num = this.stack2[len(this.stack2)-1]
		this.stack1 = append(this.stack1, num)
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}

	this.stack2 = append(this.stack2, x)
	for len(this.stack1) > 0 {
		num = this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, num)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	defer func() {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}()

	return this.stack2[len(this.stack2)-1]
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack2[len(this.stack2)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack2) == 0
}
