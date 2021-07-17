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

	for ; j < l; j++ {
		// 快指针下一位是空格,开始自转
		if sByte[j] == byte(SPACE_BYTE) {
			// slice 为长度
			revers(sByte[k:j])
			k = j + 1 // 跳过空格
		}
		if j == l-1 {
			revers(sByte[k : j+1])
		}
	}

	return string(sByte)
}

func revers(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
