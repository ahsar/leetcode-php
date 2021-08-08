package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(*testing.T) {
	var s string
	s = "cabac"
	s = "cbbd"
	s = "a"
	s = "babad"
	s = "ac"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func BenchmarkLongestPalindrome(*testing.B) {
	var s string
	s = "cabac"
	//s = "cbbd"
	//s = "a"
	r := longestPalindrome(s)
	fmt.Println(r)
}

/**
 *  最长回文串
 *
 *  1. 中心扩散法
 *
 */
func longestPalindrome(s string) string {
	var (
		length int = len(s)
		maxL   int
		l, r   int
		start  int
		tmp    int
	)

	// 1. 遍历整个字符串
	// 2. 将遍历的字符串作为中心向两边扩散
	for i := range s {
		l = i - 1
		r = i + 1
		tmp = 1

		// 左边扩散
		for l >= 0 && s[l] == s[i] {
			l--
			tmp++
		}

		// 右侧扩散
		for r < length && s[r] == s[i] {
			r++
			tmp++
		}

		// 左右两侧同时扩散
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
			tmp += 2
		}

		if tmp > maxL {
			maxL = tmp
			start = l
		}
	}
	return s[start+1 : start+maxL+1]
}
