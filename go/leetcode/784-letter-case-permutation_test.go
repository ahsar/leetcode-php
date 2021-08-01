package leetcode

import (
	"fmt"
	"testing"
	"unicode"
)

func Test784(*testing.T) {
	var s string
	s = "a1b2"
	//s = "3z4"
	//s = "12345"
	r := letterCasePermutation(s)
	fmt.Println(r)
}

func letterCasePermutation(s string) []string {
	var (
		ret  []string
		path string = ""
	)

	var backtracking func(s string, n, l int)

	backtracking = func(s string, n, l int) {
		if len(path) >= l {
			fmt.Print("---------", path, " ", n, " ")
			path = path[0 : n-2]
			fmt.Println("---------", path)
			return
		}

		for i := n; i < l; i++ {
			if unicode.IsDigit(rune(s[i])) {
				path += string(s[i])
				//fmt.Println("__", path)
				backtracking(s, n+1, l)
				return
			}

			// 放入大小写转换
			if unicode.IsUpper(rune(s[i])) {
				path += string(s[i] + 32)
				backtracking(s, n+1, l)
			} else if unicode.IsLower(rune(s[i])) {
				path += string(s[i] - 32)
				backtracking(s, n+1, l)
			}

			// 加入本体
			path += string(s[i])
			backtracking(s, n+1, l)
		}
	}

	backtracking(s, 0, len(s))
	return ret
}

// 参数: 字符串, 起始位置
