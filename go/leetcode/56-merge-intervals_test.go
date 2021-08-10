package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func Test56(*testing.T) {
	var intervals [][]int
	intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals = [][]int{{1, 4}, {4, 5}}
	//intervals = [][]int{{1, 4}, {0, 4}}
	//intervals = [][]int{{1, 4}, {2, 3}}
	//intervals = [][]int{{1, 4}, {0, 5}}
	//intervals = [][]int{{1, 4}, {0, 0}}
	//intervals = [][]int{{1, 4}, {0, 1}}
	//intervals = [][]int{{1, 4}, {0, 2}, {3, 5}}
	//intervals = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	r := merge1(intervals)
	fmt.Println(r)
}

/**
 * 合并区间
 */
func merge1(intervals [][]int) [][]int {
	var l int = len(intervals)
	if l < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var (
		res      [][]int
		resIndex int
	)

	res = append(res, intervals[0])
	// 当前数组第一项元素是否在前一项
	for i := 1; i < l; i++ {
		// 1. 当前项的1'th元素, 大于前一项2'th
		// 2. 当前项的1'th元素, 小于前一项1'th元素,同时2'th小于前一项2'th元素
		// 直接把当前项放进结果数组
		if intervals[i][0] > res[resIndex][1] {
			res = append(res, intervals[i])
			resIndex++
			continue
		}

		// 当前项2'th元素 > 前一项2'th, 前一项的右边界更新为当前项的右边界
		if intervals[i][1] > res[resIndex][1] {
			res[resIndex][1] = intervals[i][1]
		}
	}

	return res
}
