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

	//-2 -2 -2 3 3
	//gas = []int{2, 3, 4}
	//cost = []int{3, 4, 3}
	r := canCompleteCircuit(gas, cost)
	fmt.Println(r)
}

/**
 *  遍历加油站每一个节点, 对当前节点进行轮转(暴力法)
 */
func canCompleteCircuit(gas []int, cost []int) int {
	var l int = len(gas)
	var total int

	for i := range gas {
		total = 0
		for n, j := i, 0; j < l; j++ {
			// 如果不够消耗, 则外层(大循环)向后移动
			// 因为判断从当前起点是否能够轮转, 不能轮转就跳到下个节点
			if gas[(n+j)%l]+total < cost[(n+j)%l] {
				break
			}

			// 从入口处已经遍历完整一次了
			if j == l-1 {
				return i
			}
			total += gas[(n+j)%l] - cost[(n+j)%l]
		}
	}
	return -1
}
