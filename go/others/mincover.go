package others

import "fmt"

func TestMinCover() {
	var nums []int
	nums = []int{7, 5, 2, 7, 2, 7, 4, 7}
	r := minCover(nums)
	fmt.Println(r)
}

/**
 * Tesla 面试题
 * 给定一个数组, 求最小覆盖所有元素的长度
 * 提供了1 3 4 7的目的地, 最少的天数visted 所有地区
 * TODO 非正确版本
 */
func minCover(A []int) int {
	var set map[int]int = make(map[int]int)
	for _, v := range A {
		set[v] = 1
	}

	var r int = len(A)
	for i := 0; i < len(A); i++ {
		hash := make(map[int]int)
		l := 0
		for j := i; j < len(A); j++ {
			if _, ok := hash[A[j]]; !ok {
				l++
			}
			hash[A[j]] = i
		}
		if len(hash) != len(set) {
			continue
		}
		r = min(r, l)
	}
	return r
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
