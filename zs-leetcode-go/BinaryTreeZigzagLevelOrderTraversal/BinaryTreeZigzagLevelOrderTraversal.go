package BinaryTreeZigzagLevelOrderTraversal

import "for-an-offer/zs-leetcode-go/types"

func getNums(nodeQueue []*types.TreeNode, flag int) ([]int, []*types.TreeNode){
	tmp := make([]int, 0)
	tmpQueue := make([]*types.TreeNode, 0)
	start, end := 0, len(nodeQueue)
	for ; start != end; start ++{
		if flag == 1 {
			tmp = append(tmp, nodeQueue[start].Val)
		} else {
			tmp = append(tmp, nodeQueue[end-start-1].Val)
		}
		if nodeQueue[start].Left != nil {
			tmpQueue = append(tmpQueue, nodeQueue[start].Left)
		}
		if nodeQueue[start].Right != nil {
			tmpQueue = append(tmpQueue, nodeQueue[start].Right)
		}
	}
	return tmp, tmpQueue
}

func zigzagLevelOrder(root *types.TreeNode) [][]int {
	res := make([][]int, 0)
	nodeQueue := make([]*types.TreeNode, 0)
	if root == nil {
		return res
	}
	nodeQueue = append(nodeQueue, root)
	flag := 1
	for len(nodeQueue) > 0 {
		nums, queue := getNums(nodeQueue, flag)
		flag = 1-flag
		res = append(res, nums)
		nodeQueue = queue
	}
	return res
}
