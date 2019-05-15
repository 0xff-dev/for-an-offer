package ConstructBinaryTreefromInorderandPostorderTraversal

import "for-an-offer/zs-leetcode-go/types"

func buildTree(inorder []int, postorder []int) *types.TreeNode {
	inorderLen, postorderLen := len(inorder), len(postorder)
	if inorderLen == 0 {
		return nil
	}
	root := &types.TreeNode{}
	root.Val = postorder[postorderLen-1]
	// 找到
	if inorderLen == 1 {
		if postorderLen == 1 && postorder[0] == inorder[0] {
			return root
		}
		return nil
	}
	index := 0
	for ; inorder[index] != postorder[postorderLen-1]; index++ {
	}
	// 能找到
	if index > 0 {
		root.Left = buildTree(inorder[:index], postorder[:index])
	}
	if index < inorderLen {
		root.Right = buildTree(inorder[index+1:], postorder[index:postorderLen-1])
	}
	return root
}
