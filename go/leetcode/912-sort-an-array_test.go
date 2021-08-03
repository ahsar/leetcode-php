package leetcode

import (
	"fmt"
	"testing"
)

func Test912(*testing.T) {
	var nums []int
	nums = []int{5, 1, 1, 2, 0, 0}
	nums = []int{10, 3, 5, 2, 5, 1}
	nums = []int{6, 1, 3, 9, 8}

	sortArray(nums)
	fmt.Println(nums)
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partion(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

// 分区函数, 取数组最后元素, 遍历整个数组
// 小于pivot的放到左区间右顶点的下一个元素中
// 当遍历完成, 此时i则是左区间顶点+1的位置,此时交换
// 返回中间位置索引
func partion(nums []int, l, r int) int {
	privot := nums[r]

	// i 是已处理区间右边界
	// 一旦当前数值小于pivot,就把当前值放到i的右侧
	var i int = l
	for j := l; j < r; j++ {
		// 交换
		if nums[j] < privot {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}

	nums[i], nums[r] = nums[r], nums[i]
	return i
}

// 6 11 3 9 8

// i = 0
// 6 11 3 9 8
// i,j = 1

// i,j = 1
// 6 11 3 9 8
// j = 2

// i,j = 1, 2
// 6 3 11 9 8
// i,j = 2, 3

// 1 2
// i,j = 0, 0
// 1,2

// i,j,r = 2, 3, 4
// 结束循环
// 6 3 8 9 11
// return i(i 元素的两边是已经拆分的)

// 复盘:
// 关于分区函数
// 1. 摘取数组最后元素
// 2. 维护2各区间, 两个区间并不一定需要真实的数组保存, 使用指针(索引)指向范围
// 3. 遍历数组区间内元素(不要越界, 因为是递归传递, 不要直接使用数组长度)
// 4. 如果当前元素, 小于pivot, 则把最小区间的右边元素跟当前元素交换
// 5. 遍历结束, 此时i指向了左侧区间的最右位置, 从这个位置右边就是大于当前元素的区间
// 6. 返回i的索引位置

// 注意:
// 1. quickSort迭代过程中, 不要把分区中点, 也包含到排序中, 因为只针对两个子区间进行分区即可
// 2. quickSort调用一直使用索引值, 所以传入函数参数, 取到len(nums)-1
