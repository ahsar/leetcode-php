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

func combine(n int, k int) [][]int {
	var ret [][]int
	// 1 2 3 4 5 6 7 8
	ret = make([][]int, n*k)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

	return ret
}
