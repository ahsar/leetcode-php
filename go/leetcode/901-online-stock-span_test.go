// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test901(*testing.T) {
	var (
		r    int
		nums []int
	)

	// case 1
	nums = []int{100, 80, 60, 70, 60, 75, 85}

	// case 2
	//nums = []int{28, 14, 28, 35, 46, 53, 66, 80, 87, 88}

	s := Constructor1()
	for i := range nums {
		r = s.Next(nums[i])
		fmt.Println("nums:", nums[i], "res:", r)
	}
}

type StockSpanner struct {
	nums []int
	span []int
}

func Constructor1() *StockSpanner {
	return &StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	this.nums = append(this.nums, price)
	if len(this.nums) == 1 {
		// 长度等于1时, 不进行比较
		this.span = append(this.span, 1)
		return 1
	}

	// 当前项小于前一项直接1
	if price < this.nums[len(this.nums)-2] {
		this.span = append(this.span, 1)
		return 1
	} else {
		var r int
		// 不断向前遍历直到找到大于当前日期值
		for i := len(this.nums) - 2; i >= 0 && this.nums[i] <= price; {
			r += this.span[i]
			i -= this.span[i]
		}
		this.span = append(this.span, r+1)
		return r + 1
	}
}
