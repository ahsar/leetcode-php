// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"math"
	"testing"
	"unicode"
)

func TestMyAtoi(*testing.T) {
	var s string
	s = " -48"
	//s = "48"
	//s = "+48"
	s = "-111111111922337203685477595"
	s = "4193 with words"
	s = "-91283472332"
	s = "-+12"
	s = " "
	s = "21474836460"
	//s = "2147483646"
	r := myAtoi(s)
	fmt.Println("res:", r)
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	//var max int = math.MaxInt32 / 10
	var flag int = 1
	var i int
	// 清除前导空格
	for i < len(s) && string(s[i]) == " " {
		i++
	}
	if i >= len(s) {
		return 0
	}
	if string(s[i]) == "-" {
		flag = -1
	}
	if string(s[i]) == "+" || string(s[i]) == "-" {
		i++
	}

	var res int
	for ; i < len(s); i++ {
		// 非字符串直接break
		if !unicode.IsDigit(rune(s[i])) {
			break
		}

		res = res*10 + int(s[i]-'0')
		if flag == -1 && -res < math.MinInt32 {
			return math.MinInt32
		}
		if flag == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if flag > 0 {
		return res
	} else {
		return -res
	}
}
