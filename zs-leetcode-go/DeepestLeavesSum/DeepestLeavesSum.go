package DeepestLeavesSum

import "for-an-offer/zs-leetcode-go/types"

// 先计算出二叉树的最深，然后分别计算接口，这里考虑层遍历
func deepestLeavesSum(root *types.TreeNode) int {
	queue := []*types.TreeNode{root}
	res := root.Val
	for len(queue) > 0 {
		tmpQueue := make([]*types.TreeNode, 0)
		tmp := 0
		for _, item := range queue {
			tmp += item.Val
			if item.Left != nil {
				tmpQueue = append(tmpQueue, item.Left)
			}
			if item.Right != nil {
				tmpQueue = append(tmpQueue, item.Right)
			}
		}
		if len(tmpQueue) == 0 {
			res = tmp
		}
		queue = tmpQueue
	}
	return res
}
