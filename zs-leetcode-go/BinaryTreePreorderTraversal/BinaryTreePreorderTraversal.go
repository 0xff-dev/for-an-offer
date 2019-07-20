package BinaryTreePreorderTraversal

import "for-an-offer/zs-leetcode-go/types"

// 前序遍历不递归
func preorderTraversal(root *types.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*types.TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			result = append(result, root.Val)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return result
}