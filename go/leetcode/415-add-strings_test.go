package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func Test415(*testing.T) {
	var num1, num2 string
	num1 = "123"
	num2 = "456"

	r := addStrings(num1, num2)
	fmt.Println(r)
}

func addStrings(num1 string, num2 string) string {
	var (
		add int
		r   string
	)

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		tmp := x + y + add
		r = strconv.Itoa(tmp%10) + r
		add = tmp / 10
	}
	return r
}
