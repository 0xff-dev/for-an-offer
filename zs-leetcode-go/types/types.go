package types

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Display() {
	p := ln
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

type ComplexListNode struct {
	Val    int
	Next   *ComplexListNode
	Others *ComplexListNode
}

func (cln *ComplexListNode) Display() {
	walk := cln
	for walk != nil {
		fmt.Printf("val: %d others: %v", walk.Val, walk.Others)
		walk = walk.Next
	}
}

type Node struct {
	Val int
	Children []*Node
}