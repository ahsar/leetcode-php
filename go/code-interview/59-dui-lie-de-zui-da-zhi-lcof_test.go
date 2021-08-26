package codeinterview

import (
	"fmt"
	"testing"
)

func TestQue(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4}
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

	for i := len(this.deque) - 1; i > 0; i++ {
		if this.deque[i] > value {
			break
		}
		this.deque[i] = value
		this.deque[i] = 0
	}

	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	return 1
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
