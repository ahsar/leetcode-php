// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test93(*testing.T) {
	var s string
	s = "25525511135"
	s = "0000"
	s = "1111"
	s = "010010"
	s = "101023"
	r := restoreIpAddresses(s)
	fmt.Println(r)
}

func restoreIpAddresses(s string) []string {
	var res []string
	if s == "" || len(s) > 12 {
		return res
	}

	var path []string
	var backTracking = func(int) {}
	backTracking = func(start int) {
		// 终止条件
		if start >= len(s) {
			if len(path) == 4 {
				str := strings.Join(path, ".")
				if len(str)-3 == len(s) {
					res = append(res, str)
				}
			}
			return
		}

		for num, i := 0, start; i < len(s); i++ {
			num = num*10 + int(s[i]-'0')
			if num >= 0 && num <= 255 {
				path = append(path, strconv.Itoa(num))
				backTracking(i + 1)
				path = path[:len(path)-1]
			} else {
				fmt.Println("break", num)
				break
			}
		}
	}

	backTracking(0)
	return res
}
