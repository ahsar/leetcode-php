// Package  jindian provides ...
package jindian

// 回文排列

import (
	"fmt"
	"testing"
)

/**
 * 给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
 * 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
 * 回文串不一定是字典当中的单词。
 */
func Test04(*testing.T) {
	var s string

	// case1
	s = "tactcoa"

	// case2
	//s = "ttt"

	// case3
	//s = "tttt"

	r := canPermutePalindrome(s)
	fmt.Println("res:", r)
}

/**
 * 1. 统计字符串每个字符出现次数
 * 2. 如果奇数次数>=2 则不是回文串
 */
func canPermutePalindrome(s string) bool {
	var (
		nums int
		hash map[string]int = make(map[string]int)
	)

	for i := 0; i < len(s); i++ {
		hash[string(s[i])]++
	}

	for _, r := range hash {
		if (r & 1) == 0 {
			continue
		}
		nums++
		if nums >= 2 {
			return false
		}
	}
	return true
}
