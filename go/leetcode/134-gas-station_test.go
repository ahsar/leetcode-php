// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"testing"
)

func Test134(*testing.T) {
	var gas, cost []int
	gas = []int{1, 2, 3, 4, 5}
	cost = []int{3, 4, 5, 1, 2}
	canCompleteCircuit(gas, cost)
}

/**
 *  遍历加油站每一个节点, 对当前节点进行轮转(暴力法)
 */
func canCompleteCircuit(gas []int, cost []int) int {
	var l int = len(gas)
	total := 0

	for i, _ := range gas {
		for j := i; j < l; j++ {
			// 如果不够消耗, 则外层向后移动
			if gas[i]+total < cost[j%l] {
				break
			} else if j == l-1 {
				fmt.Println(gas[i], j)
			}
			total += gas[i]
			fmt.Println("--", i, j)
		}
		//fmt.Println(gas[i])
	}
	return 1
}
