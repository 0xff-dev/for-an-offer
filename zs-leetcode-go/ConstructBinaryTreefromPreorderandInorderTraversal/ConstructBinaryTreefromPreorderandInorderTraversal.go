package ConstructBinaryTreefromPreorderandInorderTraversal

import "for-an-offer/zs-leetcode-go/types"

func buildTree(preorder []int, inorder []int) *types.TreeNode {
	preorderLen, inorderLen := len(preorder), len(inorder)
	if preorderLen == 0 {
		return nil
	}
	root := &types.TreeNode{}
	root.Val = preorder[0]
	if preorderLen == 1 {
		if inorderLen == 1 && preorder[0] == inorder[0] {
			return root
		}
		return nil
	}
	index := 0
	for ; inorder[index] != preorder[0]; index ++ {}
	if index > 0 {
		root.Left = buildTree(preorder[1:index+1], inorder[:index])
	}
	if index < inorderLen {
		root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	}
	return root
}