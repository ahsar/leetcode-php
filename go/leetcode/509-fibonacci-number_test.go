package leetcode

import (
	"fmt"
	"testing"
)

func Test509(*testing.T) {
	r := fib(3)
	fmt.Println(r)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	var dp []int = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
