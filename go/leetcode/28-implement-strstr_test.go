// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test28(*testing.T) {
	var haystack, needle string

	// case 1
	haystack = "heooll"
	needle = "ll"

	// case 2
	haystack = "aaaaa"
	needle = "bba"

	// case 3
	//haystack = ""
	//needle = ""

	// case 4
	haystack = "hellc"
	needle = "lc"

	// case 5
	haystack = "lclc|"
	needle = "lc|"

	r := strStr(haystack, needle)
	fmt.Println(r)
}

/**
 * 常规版(strstr)
 *
 * 1. 挨个遍历(暴力版)
 */
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i, start := 0, 0; i < len(haystack); i, start = i+1, start+1 {
		for l, j := i, 0; l < len(haystack); {
			if haystack[l] != needle[j] {
				break
			}
			l++
			j++

			fmt.Println(string(haystack[i]), l-start, l, start, j)
			if l-start == len(needle) {
				return start
			}
		}
	}

	return -1
}
