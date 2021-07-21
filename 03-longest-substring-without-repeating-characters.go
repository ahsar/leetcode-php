package main

import (
	"fmt"
)

func main() {
	var s string

	//case 1
	s = "abcabcbb" //3 (abc)

	//case 2
	s = "bbbbb" //1 (b)

	//case 3
	s = "pwwkew" //3 (wke)

	//case 4
	s = "" //0

	//case 5
	s = "r" //1

	//case 6
	//s = "dvdf" //3

	//case 7
	s = "abba" //2

	//case 8
	s = "abcabcbb" //3

	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}

/*
 * 滑动窗口
 * 1. 用hash保存key => index
 * 2. 右指针不断向右滑动
 * 3. 右指针滑动过程中, hash遇到重复元素, 左指针跳转重复元素索引值, 同时避免case:abba, 被a重置
 * 4. 并使用max函数保存最大值
 */

// 给定一个字符串s ，请你找出其中不含有重复字符的最长子串的长度。
func lengthOfLongestSubstring(s string) int {
	var (
		l    int = len(s)
		hash     = make(map[string]int, l)
		r    int = 0
	)

	for i, j, val := 0, 0, ""; j <= l-1; j++ {
		val = string(s[j])

		if index, ok := hash[val]; ok {
			i = max(index+1, i)
		}

		hash[val] = j
		r = max(j-i+1, r)
	}

	return r
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
