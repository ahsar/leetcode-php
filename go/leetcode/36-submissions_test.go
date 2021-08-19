// Package leetcode provides ...
package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func Test36(*testing.T) {
	var raw [][]string
	var board [][]byte
	raw = [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	for i := range raw {
		board = append(board, []byte(strings.Join(raw[i], "")))
	}

	r := isValidSudoku(board)
	fmt.Println("res:", r)
}

func isValidSudoku(board [][]byte) bool {
	var row map[int]int = make(map[int]int, 9)
	var col map[int]int = make(map[int]int, 9)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[1]); j++ {
			if board[i][j] == '.' {
				fmt.Println("-------", board[i][j])
				continue
			}
			fmt.Println(board[i][j], int(board[i][j]))
			//break
			row[int(i)] = col[int(board[i][j])]
		}
		break
	}
	fmt.Println(board)
	var r bool
	return r
}
