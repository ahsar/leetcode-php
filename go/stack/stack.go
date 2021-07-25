package stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(v interface{}) {
	stack.list.PushBack(v)
}

func (stack *Stack) Pop() interface{} {
	if stack.Empty() {
		return nil
	}
	e := stack.list.Back()
	stack.list.Remove(e)
	return e.Value
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() <= 0
}
