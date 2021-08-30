// Package jindian provides ...
package jindian

/**
 * 栈的最小值
 */
import (
	"fmt"
	"testing"
)

func Test0302(*testing.T) {
	var (
		r   int
		obj MinStack
	)

	obj = Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(2)
	obj.Push(-3)
	obj.Push(-4)
	obj.Push(3)

	obj.Pop()
	obj.Pop()
	obj.Pop()
	r = obj.GetMin()
	fmt.Println(r)
}

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) <= 0 {
		this.minStack = append(this.minStack, x)
	} else {
		t := this.minStack[len(this.minStack)-1]
		this.minStack = append(this.minStack, min(x, t))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
