package FullBinaryTree

import "for-an-offer/zs-leetcode-go/types"

func IsFullBinaryTree(tree *types.TreeNode) bool {
	if tree == nil {
		return true
	}
	// 条件就是这个节点两个都是nil, nil 或者d
	if tree.Right == nil && tree.Left == nil {
		return true
	}
	if tree.Left != nil && tree.Right != nil {
		left := IsFullBinaryTree(tree.Left)
		right := IsFullBinaryTree(tree.Right)
		return left && right
	}
	return false
}
