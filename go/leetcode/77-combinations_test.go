package leetcode

import (
	"fmt"
	"testing"
)

func Test77(*testing.T) {
	var (
		n, k int
	)

	//case 1
	n = 8
	k = 3

	r := combine(n, k)
	fmt.Println(r)
}

var ret [][]int

type Path struct {
	v []int
}

// go 实现栈
// https://www.cnblogs.com/TimLiuDream/p/9902496.html
func (p *Path) Push(v interface{}) {
	//p.v
}

func combine(n int, k int) [][]int {
	backtracking(n, k, 1)
	return ret
}

func backtracking(n, k, i int) {
}
