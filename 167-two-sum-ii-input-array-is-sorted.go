package main

/**
 * 给定一个已按照 升序排列
 * 的整数数组 numbers,
 * 请你从数组中找出两个数满足相加之和等于目标数 target 。
 */
import "fmt"

func main() {
	var (
		nums, r []int
		target  int
	)

	// case 1
	nums = []int{2, 7, 11, 15}
	target = 9

	// case 2
	//nums = []int{2, 3, 4}
	//target = 6

	// case 3
	//nums = []int{-1, 0}
	//target = -1

	r = twoSum(nums, target)
	fmt.Println(r)
}

func twoSum(numbers []int, target int) []int {
	var (
		l, r, sum int = 0, len(numbers) - 1, 0
		ret       []int
	)
	ret = make([]int, 2, 2)

	for l < r {
		sum = numbers[l] + numbers[r]

		if sum == target {
			ret[0] = l + 1
			ret[1] = r + 1
			return ret
		}

		// 右指针左移
		if sum > target {
			r--
		} else {
			l++
		}
	}
	return ret
}
