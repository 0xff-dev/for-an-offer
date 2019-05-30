package BalancedBinaryTree

import "for-an-offer/zs-leetcode-go/types"

func aux(root *types.TreeNode, res *bool) int {
	if root == nil {
		return 0
	}
	left := aux(root.Left, res)+1
	right := aux(root.Right, res)+1
	dep := left-right
	if !(dep >= -1 && dep <= 1){
		*res = false
		return 0
	}
	if dep == -1 {
		return right
	}
	return left
}
func isBalanced(root *types.TreeNode) bool {
	res := true
	aux(root, &res)
	return res
}
