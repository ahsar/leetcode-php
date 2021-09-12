// Package leetcode provides ...
package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func Test703(*testing.T) {
	var (
		nums   []int
		k      int
		stream []int
	)

	// case 1
	nums = []int{4, 5, 8, 2}
	k = 3
	stream = []int{3, 5, 10, 9, 4}

	obj := Constructor703(k, nums)
	for _, r := range stream {
		tmp := obj.Add(r)
		fmt.Println("k'th num is:", tmp)
	}
}

// 优先队列
type KthLargest struct {
	sort.IntSlice
	k int
}

/**
 * 数据流中的第 K 大元素
 */
func Constructor703(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, r := range nums {
		kl.Add(r)
	}
	return kl
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	q := this.IntSlice
	defer func() {
		this.IntSlice = q[:len(q)-1]
	}()

	return q[len(q)-1]
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}

	fmt.Println(this.IntSlice)
	return this.IntSlice[0]
}
