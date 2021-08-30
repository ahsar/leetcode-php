package jianzhi

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func Test569(*testing.T) {
	var timePoints []string

	// case 1
	timePoints = []string{"23:59", "00:00"}
	// "00:00" "23:59"
	// "00:01" 00:04 "23:59"

	// case 2
	//timePoints = []string{"00:00", "23:59", "00:00"}

	r := findMinDifference(timePoints)
	fmt.Println("res:", r)
}

/**
 * 最小时间差
 *
 * 1. 转换为分钟
 * 2. 排序
 * 3. 将数组首元素*24*60 放到元素末尾
 * 4. 作差
 *
 */
func findMinDifference(timePoints []string) int {
	if len(timePoints) > 24*60 {
		return 0
	}
	var r int = 12 * 60
	var timeList []int

	for i := 0; i < len(timePoints); i++ {
		tmp := timePoints[i]
		hour, _ := strconv.Atoi(tmp[:2])
		minute, _ := strconv.Atoi(tmp[3:])
		timeList = append(timeList, hour*60+minute)
	}

	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	sort.Ints(timeList)
	timeList = append(timeList, timeList[0]+24*60)
	for i := 1; i < len(timeList); i++ {
		if timeList[i-1] == timeList[i] {
			return 0
		}

		r = min(r, abs(timeList[i-1]-timeList[i]))
	}

	return r
}
