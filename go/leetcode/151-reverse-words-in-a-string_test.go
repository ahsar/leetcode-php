// Package leetcode provides ...
package leetcode

/**
 * 翻转字符串里的单词
 *
 * 给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
 * 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
 * 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
 * 说明：
 * 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
 * 翻转后单词间应当仅用一个空格分隔。
 * 翻转后的字符串中不应包含额外的空格。
 */

import (
	"fmt"
	"testing"
)

func Test151(*testing.T) {
	var s string
	s = "the sky is blue"
	//s = " hello world "
	//s = "a good   example"
	s = "  Bob    Loves  Alice   "
	s = "Alice does not even like bob"
	s = "a"
	r := reverseWords(s)
	fmt.Print("res:{", r, "}\n")
}

/**
 * 思路
 *
 * 遍历原始字符串
 *  定义临时中间变量并不断拼接字符串
 * 	遇到空格则把当前字符串索引直接拼接到栈中
 *  遍历完成弹栈
 */
func reverseWords(s string) (res string) {
	var (
		length            int = len(s)
		start, stackIndex int
		stack             []string = make([]string, length)
	)

	push := func(s string) {
		if s == "" {
			return
		}
		fmt.Print("push:{", s, "}\n")
		stack[stackIndex] = s
		stackIndex++
	}

	for i, l := 0, 0; i < length; i++ {
		if s[i] == 32 || i == length-1 {
			if i == length-1 && s[i] != 32 {
				push(s[start : i-l+1])
			} else {
				push(s[start : i-l])
			}
			// 到达空格处需要把单词放入栈中
			start = i + 1
			l = 0
		}
	}

	stackIndex--
	for stackIndex >= 0 {
		if stackIndex == 0 {
			res += stack[stackIndex]
		} else {
			res += stack[stackIndex] + " "
		}
		stackIndex--
	}

	return
}
