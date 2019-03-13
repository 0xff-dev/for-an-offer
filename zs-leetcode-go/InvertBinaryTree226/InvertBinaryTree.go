package InvertBinaryTree226

import "for-an-offer/zs-leetcode-go/types"

func InvertBinaryTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return  nil
	}
	var innerFunc func(root *types.TreeNode)
	innerFunc = func(root *types.TreeNode) {
		root.Left, root.Right = root.Right, root.Left
		if root.Left != nil {
			innerFunc(root.Left)
		}
		if root.Right != nil {
			innerFunc(root.Right)
		}
	}
	innerFunc(root)
	return root
}
