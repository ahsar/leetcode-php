package leetcode

import (
	. "algo/bintree"
	"testing"
)

func Test116(t *testing.T) {
	var tree []int

	tree = []int{1, 2, 3, 4, 5, 6, 7}
	var n Node
	r := n.BuildTree(0, len(tree), tree)
	//n.LevelOrder(r)
	//fmt.Println("----")
	connect(r)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

//   1
// 2   3
//4 5 6 7
func connect(root *Node) *Node {
	if nil == root {
		return nil
	}

	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = nil

		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
	}

	return root
}
