package types

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 树的层遍历
func (tree *TreeNode) Floor() {
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		for _, item := range queue {
			fmt.Print(item.Val, " ")
			if item.Left != nil {
				tmpQueue = append(tmpQueue, item.Left)
			}
			if item.Right != nil {
				tmpQueue = append(tmpQueue, item.Right)
			}
		}
		fmt.Println()
		queue = tmpQueue
	}
}