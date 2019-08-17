package BinaryTreeRightSideView

import "for-an-offer/zs-leetcode-go/types"

func rightSideView(root *types.TreeNode) []int {
	queue := make([]*types.TreeNode, 0)
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	result = append(result, root.Val)
	for len(queue) > 0 {
		tmp := make([]*types.TreeNode, 0)
		var lastNode *types.TreeNode
		for _, item := range queue {
			if item.Left != nil {
				lastNode = item.Left
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				lastNode = item.Right
				tmp = append(tmp, item.Right)
			}
		}
		if lastNode != nil {
			result = append(result, lastNode.Val)
		}
		queue = tmp
	}
	return result
}
