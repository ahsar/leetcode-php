package main

import (
	"fmt"
	"testing"
)

var stack []int
var stack1 []int

func TestStack(*testing.T) {
	var nums []int
	nums = []int{1, 2, 3, 4}
	for i := 0; i < len(nums); i++ {
		push(nums[i])
	}

	for i := 0; i < len(nums)-1; i++ {
		r := pop()
		fmt.Println("pop:", r)
	}
	push(5)
	r := pop()
	fmt.Println("pop:", r)
}

// stack1
// stack2
// 1 2 3 4 5
func push(i int) {
	stack1 = append(stack1, i)
}

func pop() int {
	for len(stack1) > 0 {
		top := stack1[len(stack1)-1]
		stack = append(stack, top)
		stack1 = stack1[:len(stack1)-1]
	}

	defer func() {
		stack = stack[:len(stack)-1]
	}()

	return stack[len(stack)-1]
}
