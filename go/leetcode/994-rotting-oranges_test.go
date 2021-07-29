package leetcode

import (
	"fmt"
	"testing"
)

func Test994(*testing.T) {
	var grid [][]int
	// case 1
	// 2 1 1
	// 1 1 0
	// 0 1 1
	grid = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}

	// case 2
	// [2,1,1]
	// [0,1,1]
	// [1,0,1]
	grid = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}

	// case 3
	grid = [][]int{{0, 2}}

	// case 4
	grid = [][]int{{1, 2}}

	// case 5
	grid = [][]int{{1}}

	// case 6
	grid = [][]int{{1, 0}}

	// case 7
	grid = [][]int{{0}}
	r := orangesRotting(grid)
	fmt.Println(r)
}

// BFS(广度优先遍历)
// 1. 遍历网格
// 2. 遇到橘子+1,遇到腐烂橘子, 加入队列
// 3. while(队列不为空)
// 4. 	一次性取出队列所有元素
// 5. 	把腐烂橘子的四个方向依次加入队列(避免1,2 这种情况)
// 6. 返回while 次数(如果网格中事先存在橘子), 否则返回0
func orangesRotting(grid [][]int) int {
	var (
		r, c               int = len(grid), len(grid[0])
		tmp                [2]int
		que                [][2]int
		nums               int
		note               [][]int = make([][]int, r) // 避免重复添加
		flag               bool    = false
		orangNums, badNums int
	)

	// 初始化备忘录
	for i := range note {
		note[i] = make([]int, c)
	}

	// 将腐烂橘子入队列
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// 总共夺少个橘子
			if grid[i][j] == 2 || grid[i][j] == 1 {
				flag = true
				orangNums++
			}
			if grid[i][j] == 2 {
				tmp[0] = i
				tmp[1] = j
				note[i][j] = 1
				badNums++
				que = append(que, tmp)
			}
		}
	}

	for len(que) > 0 {
		que2 := que
		que = nil
		for _, v := range que2 {
			// 向下
			if v[0] < r-1 && grid[v[0]+1][v[1]] == 1 && note[v[0]+1][v[1]] != 1 {
				tmp[0] = v[0] + 1
				tmp[1] = v[1]
				que = append(que, tmp)
				note[v[0]+1][v[1]] = 1
				badNums++
			}

			// 向右
			if v[1] < c-1 && grid[v[0]][v[1]+1] == 1 && note[v[0]][v[1]+1] != 1 {
				tmp[0] = v[0]
				tmp[1] = v[1] + 1
				que = append(que, tmp)
				note[v[0]][v[1]+1] = 1
				badNums++
			}

			// 向上
			if v[0] > 0 && grid[v[0]-1][v[1]] == 1 && note[v[0]-1][v[1]] != 1 {
				tmp[0] = v[0] - 1
				tmp[1] = v[1]
				que = append(que, tmp)
				note[v[0]-1][v[1]] = 1
				badNums++
			}

			// 向左
			if v[1] > 0 && grid[v[0]][v[1]-1] == 1 && note[v[0]][v[1]-1] != 1 {
				tmp[0] = v[0]
				tmp[1] = v[1] - 1
				que = append(que, tmp)
				note[v[0]][v[1]-1] = 1
				badNums++
			}
		}

		nums++
	}

	// 存在新鲜的橘子
	if badNums < orangNums && flag {
		return -1
	}
	if flag {
		nums--
	}
	return nums
}
