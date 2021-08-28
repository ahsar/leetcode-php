package jindian

// 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

import (
	"fmt"
	"reflect"
	"testing"
)

func Test02(*testing.T) {
	var (
		s1 string
		s2 string
	)

	// case 1
	s1 = "abc"
	s2 = "cba"

	// case 2
	s1 = "abc"
	s2 = "bad"

	// case 3
	s1 = " "
	s2 = " "

	r := CheckPermutation(s1, s2)
	fmt.Println("res:", r)
}

func CheckPermutation(s1 string, s2 string) bool {
	var (
		hash1 map[string]int = make(map[string]int, len(s1))
		hash2 map[string]int = make(map[string]int, len(s2))
	)

	for i := 0; i < len(s1); i++ {
		hash1[string(s1[i])]++
	}

	for i := 0; i < len(s2); i++ {
		hash2[string(s2[i])]++
	}
	if reflect.DeepEqual(hash1, hash2) {
		return true
	}
	return false
}
