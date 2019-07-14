package BinaryTreeLevelOrderTraversalII

import "for-an-offer/zs-leetcode-go/types"

func aux(nodes []*types.TreeNode, res *[][]int) {
	tmp := make([]*types.TreeNode, 0)
	nums := make([]int, 0)
	for _, node := range nodes {
		nums = append(nums, node.Val)
		if node.Left != nil {
			tmp = append(tmp, node.Left)
		}
		if node.Right != nil {
			tmp = append(tmp, node.Right)
		}
	}
	if len(tmp) != 0 {
		aux(tmp, res)
	}
	*res = append(*res, nums)
}

func levelOrderBottom(root *types.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	nodes := []*types.TreeNode{root}
	aux(nodes, &res)
	return res
}