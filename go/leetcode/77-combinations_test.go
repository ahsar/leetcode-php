package leetcode

import (
	"fmt"
	"testing"
)

// 需要理解深度优先遍历
// 深度优先遍历, 利用递归, 使其停留在根节点. 然后在根节点进行遍历
// 深度优先遍历 三部曲 1. 递归函数参数 2. 终止条件 3. 循环
func Test77(*testing.T) {
	var (
		n, k int
	)

	//case 1
	n = 1
	k = 1

	r := combine(n, k)
	fmt.Println(r)
}

var (
	path []int
)

// go 实现栈
// https://www.cnblogs.com/TimLiuDream/p/9902496.html
func push(v int) {
	path = append(path, v)
}

func pop() int {
	if len(path) <= 0 {
		return 0
	}

	v := path[len(path)-1]
	path = path[:len(path)-1]
	return v
}

func combine(n int, k int) [][]int {
	result := [][]int{}

	var backtracking func(n, k, start int)
	backtracking = func(n, k, start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := start; i <= n; i++ {
			push(i)
			backtracking(n, k, i+1) // 1 进入栈后, 随着叶子节点的执行完毕. 1也被pop
			pop()
		}
	}

	backtracking(n, k, 1)
	return result
}
