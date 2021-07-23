package leetcode

import (
	"fmt"
	"testing"
)

/*
 *给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 *换句话说，第一个字符串的排列之一是第二个字符串的子串 。
 */
func Test567(t *testing.T) {
	var (
		s1, s2 string
		r, ans bool
	)

	//case 1
	s1 = "ab"
	s2 = "eidbaooo"
	ans = true

	//case 2
	s1 = "ab"
	s2 = "eidboaoo"
	ans = false

	//case 3
	s1 = "abc"
	s2 = "eidcbaoo"
	ans = true

	//case 4
	s1 = "ab"
	s2 = "ab"
	ans = true

	//case 5
	s1 = "hello"
	s2 = "ooolleoooleh"
	ans = false

	// case 6
	s1 = "adc"
	s2 = "dcda"
	ans = true

	if r = checkInclusion(s1, s2); ans != r {
		t.Error("case not pass ans: ", r)
	}
	fmt.Println("res is:", r)
}

// 滑动窗口(双指针)
// 约束: s1,s2 >=1
//
// 存在缺陷
// 1. 用hash存储s1 每个元素
// 2. 遍历s2, 一旦匹配右指针继续向前滑动, 否则左指针=当前节点的下一个节点
// 重复使用元素问题. 一旦不匹配, 无法清理上一次因为半匹配(多个) 的数量
// 3. 匹配后, 判断窗口长度是否等于s1长度
//
// 正确方法
// 思路: 对比map方式
// 1. 针对s1 创建map[string]int, 保存出现次数
// 2. 遍历s2 创建map[string]int, 当窗口等于s1长度, 判断两者map是否相等
// 3. 窗口大小不变,向右平移,左指针自增, 同时,map的左指针--, 保证元素数量
//
// ticket 因为只有26个英文字符, 所以可以使用数组作为map, 索引为字母, v是数量
func checkInclusion(s1 string, s2 string) bool {
	var (
		l1          = len(s1)
		l2          = len(s2)
		hash, hash2 [26]int
	)

	if l1 > l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		hash[s1[i]-'a']++
	}

	// todo 可以使用l1, 代替i
	for i, j := 0, 0; j < l2; j++ {
		// 对s2 进行hash
		hash2[s2[j]-'a']++

		//s2 = "eidcbaoo"
		if j-i+1 >= l1 {
			if hash == hash2 {
				return true
			}
			hash2[s2[i]-'a']-- // 左指针字符串弹出
			i++
		}
	}
	return false
}
