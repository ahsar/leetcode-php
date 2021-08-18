// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test763(*testing.T) {
	var s string
	s = "ababcbacadefegdehijhklij"
	s = "aba"
	r := partitionLabels(s)
	fmt.Println(r)
}

func partitionLabels(s string) []int {
	var res []int

	// 构建映射hash
	var hash map[string]int = make(map[string]int)
	for i := 0; i < len(s); i++ {
		hash[string(s[i])] = max(hash[string(s[i])], i)
	}

	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	// 1. 遍历字符串获取当前字符串的最后位置
	// 2. 维护start, end, end 保存当前字符串的最后位置
	// 3. 如果索引抵达最后位置, 则当前片段访问结束
	for i, start, end := 0, 0, 0; i < len(s); i++ {
		end = max(end, hash[string(s[i])])
		if i >= end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
