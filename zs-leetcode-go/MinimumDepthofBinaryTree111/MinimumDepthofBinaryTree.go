package MinimumDepthofBinaryTree111

import "for-an-offer/zs-leetcode-go/types"

func MinimumDepthofBinaryTree(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	left := MinimumDepthofBinaryTree(root.Left)
	right := MinimumDepthofBinaryTree(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}
