// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test213(*testing.T) {
	var n int
	n = -2
	r := isPowerOfTwo(n)
	fmt.Println(r)
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
