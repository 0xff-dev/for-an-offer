package KthSmallestElementInABST

import "for-an-offer/zs-leetcode-go/types"

func kthSmallest(root *types.TreeNode, k int) int {
	if root == nil {
		return -1
	}
	var res int = -1
	now := 0
	inOrder(root, k, &now, &res)
	return res
}

func inOrder(root *types.TreeNode, kth int,now, val *int) {
	if root.Left != nil {
		inOrder(root.Left,kth, now, val)
	}
	*now = *now + 1
	if *now == kth {
		*val = root.Val
		return
	}
	if root.Right != nil {
		inOrder(root.Right, kth, now, val)
	}
}
