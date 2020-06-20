package InsufficientNodesinRoottoLeafPaths

import (
	"for-an-offer/zs-leetcode-go/types"
)

func sufficientSubset(root *types.TreeNode, limit int) *types.TreeNode {
	if root == nil {
		return nil
	}
	return aux(root, 0, limit)
}

func aux(root *types.TreeNode, now, limit int) *types.TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		now += root.Val
		if now < limit {
			root = nil
		}
		return root
	}
	root.Left = aux(root.Left, now + root.Val, limit)
	root.Right = aux(root.Right, now + root.Val, limit)
	if root.Right == nil && root.Left == nil {
		return nil
	}
	return root
}
