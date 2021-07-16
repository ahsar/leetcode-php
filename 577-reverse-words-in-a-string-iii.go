package main

import (
	"fmt"
)

const SPACE_BYTE int = 32

func main() {
	var s string
	s = "Let's take LeetCode contest" // s'teL ekat edoCteeL tsetnoc
	r := reverseWords(s)
	fmt.Println(r)
}

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
func reverseWords(s string) string {
	var (
		sByte     = []byte(s)
		l     int = len(sByte)
		j, k  int = 0, 0
	)

	for i := 0; i < l; i++ {
		// 快指针下一位是空格,开始自转
		if sByte[j+1] == byte(SPACE_BYTE) {
			// slice 为长度
			revers(sByte[k : j-1])
			j++
			k += j + 1
			//break
		}
		j++
	}

	return string(sByte)
}

func revers(s []byte) {
	fmt.Println(string(s))
	var (
		l int = len(s)
	)
	mid := l / 2

	for i := 0; i < mid; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
