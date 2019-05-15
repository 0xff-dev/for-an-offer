package BinaryTreeLevelOrderTraversal

import "for-an-offer/zs-leetcode-go/types"

func levelOrder(root *types.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*types.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]*types.TreeNode, 0)
		tmpNum := make([]int, 0)
		for _, item := range queue {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
			tmpNum = append(tmpNum, item.Val)
		}
		res = append(res, tmpNum)
		queue = tmp
	}
	return res
}
