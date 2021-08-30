package codeinterview

import (
	"fmt"
	"testing"
)

func TestQue(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4, 1, 2}
	que := Constructor()

	for _, v := range nums {
		que.Push_back(v)
	}
	fmt.Println(que.deque, que.que)
}

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
	this.deque = append(this.deque, value)
	if len(this.deque) == 0 {
		return
	}

	// 1 2 3 4 | 1 2
	// 4 2
	var i int
	for i = len(this.deque) - 1; i > 0; i-- {
		if this.deque[i-1] > value {
			break
		}
		this.deque[i-1] = this.deque[i]
	}
	// 4
	// 7 5 3 1

	fmt.Println(this.deque)
	//this.deque = append(this.deque, value)
	//this.deque[i] = value
}

func (this *MaxQueue) Pop_front() int {
	return 1
}
