package leetcode

import (
	"fmt"
	"testing"
)

func TestCQueue(*testing.T) {
	cq := Constructor2()
	for i := 1; i < 5; i++ {
		fmt.Println("in:", i)
		cq.AppendTail(i)
	}

	var r int
	for i := 1; i < 2; i++ {
		r = cq.DeleteHead()
		fmt.Println("out", r)
	}

	cq.AppendTail(7)
	r = cq.DeleteHead()
	fmt.Println("out", r)
}

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor2() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	// stack2 不为空, 不断弹到stack1
	for len(this.stack2) > 0 {
		tmp := this.stack2[len(this.stack2)-1]
		this.stack1 = append(this.stack1, tmp)
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	// stack2 为空
	this.stack2 = append(this.stack2, value)

	for len(this.stack1) > 0 {
		tmp := this.stack1[len(this.stack1)-1]
		this.stack2 = append(this.stack2, tmp)
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) <= 0 {
		return -1
	}
	defer func() {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}()

	return this.stack2[len(this.stack2)-1]
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
