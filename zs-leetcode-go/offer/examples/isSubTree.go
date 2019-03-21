package examples

import (
	"for-an-offer/zs-leetcode-go/types"
)

func match(tree1, tree2 *types.TreeNode) bool {
	if tree1 == nil {
		return false
	}
	if tree2 == nil {
		return true
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return match(tree1.Left, tree1.Left) && match(tree1.Right, tree2.Right)
}

func IsSubTree(tree1, tree2 *types.TreeNode) bool {
	res := false
	if tree1 != nil && tree2 != nil {
		if tree1.Val == tree2.Val {
			res = match(tree1, tree2)
		}
		if !res {
			res = IsSubTree(tree1.Left, tree2)
		}
		if !res {
			res = IsSubTree(tree1.Right, tree2)
		}
	}
	return res
}
