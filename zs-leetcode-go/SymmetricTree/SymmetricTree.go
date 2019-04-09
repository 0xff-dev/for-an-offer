package SymmetricTree

import "for-an-offer/zs-leetcode-go/types"

func isSymmetric(root *types.TreeNode) bool {
	if root == nil {
		return false
	}
	return auxFunc(root.Left, root.Right)
}

func auxFunc(node1, node2 *types.TreeNode) bool {
	if node1 == nil && node2 != nil || node1 != nil && node2 == nil {
		return false
	}
	if node1 == nil && node2 == nil {
		return true
	}
	if node1.Val == node2.Val {
		return auxFunc(node1.Left, node2.Right) && auxFunc(node1.Right, node2.Left)
	}
	return false
}
