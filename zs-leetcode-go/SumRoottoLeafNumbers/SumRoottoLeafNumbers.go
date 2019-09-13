package SumRoottoLeafNumbers

import "for-an-offer/zs-leetcode-go/types"

func aux(node *types.TreeNode, tmp int, sum *int) {
	flag := false
	if node.Left != nil {
		aux(node.Left, tmp*10+node.Left.Val, sum)
		flag = true
	}
	if node.Right != nil {
		aux(node.Right, tmp*10+node.Right.Val, sum)
		flag = true
	}
	if !flag {
		*sum += tmp
	}
}

func sumNumbers(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	aux(root, root.Val, &result)
	return result
}
