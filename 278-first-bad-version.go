// Package main provides ...
// 二分法查找
package main

import "fmt"

var bad int

func main() {
	var n int

	n = 5
	bad = 4 // case 1

	n = 1
	bad = 1 // case 2

	n = 4
	bad = 1 // case 3

	r := firstBadVersion(n)
	fmt.Println(r)
}

func firstBadVersion(n int) int {
	var (
		l, mid int
	)
	l = 0

	for l < n {
		mid = (l + n) / 2
		if isBadVersion(mid) {
			// 找到第一个错误版本
			n = mid
		} else {
			// 否则右移
			l = mid + 1
		}
	}

	return n
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(n int) bool {
	if n >= bad {
		return true
	}
	return false
}
