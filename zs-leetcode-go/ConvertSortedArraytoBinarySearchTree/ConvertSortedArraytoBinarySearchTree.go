package ConvertSortedArraytoBinarySearchTree

import "for-an-offer/zs-leetcode-go/types"

// 构建左右子树，flag是判断左右子树的条件
func aux(root *types.TreeNode, flag bool, datas []int) {
	if len(datas) == 0 {
		return
	}
	mid := len(datas) / 2
	node := &types.TreeNode{datas[mid], nil, nil}
	aux(node, true, datas[0:mid])
	aux(node, false, datas[mid+1:])
	if flag {
		root.Left = node
	} else {
		root.Right = node
	}
}

func sortedArrayToBST(nums []int) *types.TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	mid := length / 2
	root := &types.TreeNode{nums[mid], nil, nil}
	aux(root, true, nums[:mid])
	aux(root, false, nums[mid+1:])
	return root
}
