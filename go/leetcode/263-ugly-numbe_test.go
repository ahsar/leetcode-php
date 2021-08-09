package leetcode

import (
	"fmt"
	"testing"
)

func Test263(*testing.T) {
	var n int
	n = 6
	//n = 8
	//n = 14
	//n = 1
	r := isUgly(n)
	fmt.Println(r)
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for (n % 2) == 0 {
		n /= 2
	}
	for (n % 5) == 0 {
		n /= 5
	}
	for (n % 3) == 0 {
		n /= 3
	}

	if n == 1 {
		return true
	}
	return false
}
