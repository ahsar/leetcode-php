// Package jindian provides ...
package jindian

// 字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成(比如，waterbottle是erbottlewat旋转后的字符串)

import (
	"fmt"
	"testing"
)

func Test09(*testing.T) {
	var (
		s1, s2 string
	)

	// case 1
	s1 = "waterbottle"
	s2 = "erbottlewat"

	// case 2
	s1 = "aa"
	s2 = "aba"

	// case 3
	//s1 = "abac"
	//s2 = "abad"

	// case 4
	//s1 = "PvcvpkpHwaXQxpgGzURBvHRMvCsCPPmlKBSzXDWSvrxLBPdAvRpgcIwNOVQDdwPIElrAFqmb"
	//s2 = "SvrxLBPdAvRpgcIwNOVQDdwPIElrAFqmbPvcvpkpHwaXQxpgGzURBvHRMvCsCPPmlKBSzXDW"

	// case 5
	s1 = ""
	s2 = ""

	r := isFlipedString(s1, s2)
	fmt.Println("res:", r)

}

/**
 * 1. s1 拼接s1
 * 2. 遍历s1
 * 3. 若全匹配到s2, 则 s2 是 s1 的旋转
 */
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == "" {
		return true
	}

	s1 += s1
	var k int
	for i := 0; i < len(s1); i++ {
		k = 0

		if s1[i] == s2[k] {
			for j := i; j < len(s1) && k < len(s2) && s1[j] == s2[k]; j, k = j+1, k+1 {
			}

			//fmt.Println(j, k, 1)
			if k >= len(s2) {
				return true
			}
		}
	}
	return false
}
