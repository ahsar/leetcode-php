package leetcode

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
import (
	"fmt"
	"testing"
)

func Test46(*testing.T) {
	var nums []int

	// case 1
	nums = []int{1, 2, 3}

	// case 2
	//nums = []int{0, 1}

	// case 3
	//nums = []int{1}

	// case 4
	//nums = []int{1, 8, 9, 12}
	r := permute(nums)
	fmt.Println(r)
}

func permute(nums []int) [][]int {
	var (
		ret          [][]int
		l            int         = len(nums)
		visted       map[int]int = make(map[int]int, l)
		backtracking func(nums []int, l int)
	)

	backtracking = func(nums []int, l int) {
		if len(visted) == l {
			tmp := make([]int, l)
			copy(tmp, path)
			ret = append(ret, tmp)
			return
		}

		for i := 0; i < l; i++ {
			// 存在
			if _, ok := visted[nums[i]]; ok {
				continue
			}

			visted[nums[i]] = i
			push(nums[i])
			backtracking(nums, l)
			delete(visted, nums[i])
			pop()
		}
	}

	backtracking(nums, l)
	return ret
}
