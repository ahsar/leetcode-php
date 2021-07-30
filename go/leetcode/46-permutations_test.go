package leetcode

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
import (
	"testing"
)

func Test46(*testing.T) {
	var nums []int
	// 1, 2 ,3
	// 1, 3, 2
	// 2, 1, 3
	// 2, 3, 1
	// 3, 1, 2
	// 3, 2, 1

	nums = []int{1, 2, 3}
	permute(nums)
}

func permute(nums []int) [][]int {
	var ret [][]int
	var l int = len(nums)
	var visted map[int]int = make(map[int]int, l)

	// 1 2 3
	// 1 3 2

	// 3 2 1
	// 3 1 2

	// 2 3 1
	// 2 1 3
	backtracking := func(nums []int, l, now int) {
		if len(path) == l {
			tmp := make([]int, l)
			copy(tmp, path)
			ret = append(ret, tmp)
			return
		}

		for i := now; i < l; i++ {
			//if {

			//}
		}
	}
	backtracking(nums, l, 0)

	return ret
}
