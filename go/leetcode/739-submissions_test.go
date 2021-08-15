package leetcode

import (
	"fmt"
	"testing"
)

func Test739(*testing.T) {
	var temperatures []int
	// case 1
	temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}

	// case 2
	temperatures = []int{30, 40, 50, 60}

	// case 3
	temperatures = []int{30, 60, 90}

	// case 4
	temperatures = []int{30}

	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}

/**
 * 每日温度
 *
 * (单调栈实现)
 * 1. 遍历温度数组
 * 2. 元素依次入栈
 * 3. 如果当前元素大于栈顶元素, 则依次弹栈并且更新索引数, 直到栈为空
 *
 */
func dailyTemperatures(temperatures []int) []int {
	var (
		l, size int = len(temperatures), 0
		stack   []int
		res     []int = make([]int, l)
	)

	for i := 0; i < l; i++ {
		size = len(stack)
		if size == 0 {
			// 放置索引
			stack = append(stack, i)
			continue
		}

		// 当前元素大于栈顶, 弹栈&& 栈内元素小于当前元素 更新结果集
		if temperatures[i] > temperatures[stack[size-1]] {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
				res[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	return res
}
