package leetcode

import (
	"fmt"
	"testing"
	"unicode"
)

func Test784(*testing.T) {
	var s string
	s = "a1b2"
	s = "3z4"
	s = "12345"
	r := letterCasePermutation(s)
	fmt.Println(r)
}

func letterCasePermutation(s string) []string {
	var ret []string

	var backtracking func(path string, i, l int)
	backtracking = func(path string, i, l int) {
		if len(path) >= l {
			ret = append(ret, path)
			return
		}

		// 大小写转换
		if unicode.IsLetter(rune(s[i])) {
			backtracking(path+string(s[i]^32), i+1, l)
		}
		// 加入本体
		backtracking(path+string(s[i]), i+1, l)
	}

	backtracking("", 0, len(s))
	return ret
}
