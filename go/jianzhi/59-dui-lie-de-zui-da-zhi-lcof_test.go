package jianzhi

import (
	"fmt"
	"testing"
)

func TestQue(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 11, 1, 2, 5, 5, 4, 3, 2}
	que := Constructor()

	for _, v := range nums {
		que.Push_back(v)
		fmt.Println("push:", v, que.deque)
	}

	var r, max int
	for range nums {
		r = que.Pop_front()
		max = que.Max_value()
		fmt.Println("pop", r, "max:", max)
	}
}

/**
 * 队列的最大值
 * 实现原理(双端队列)
 * 一个队列保存元素
 * 另一个队列保存最大值遇到最大值则进行追加
 */
type MaxQueue struct {
	que   []int
	deque []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) <= 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.que = append(this.que, value)
	// 双端队列为空直接插入
	if len(this.deque) == 0 {
		this.deque = append(this.deque, value)
		return
	}
	// 元素小于双端队列末尾元素直接追加
	if value < this.deque[len(this.deque)-1] {
		this.deque = append(this.deque, value)
		return
	}
	if value == this.deque[len(this.deque)-1] {
		return
	}

	for len(this.deque) > 0 && value > this.deque[len(this.deque)-1] {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.deque) <= 0 {
		return -1
	}

	var x int
	x = this.que[0]
	this.que = this.que[1:]

	if this.deque[0] == x {
		this.deque = this.deque[1:]
	}

	return x
}
