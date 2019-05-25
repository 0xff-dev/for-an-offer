package PathSum

import "for-an-offer/zs-leetcode-go/types"

func hasPathAux(root *types.TreeNode, now, target int, res *bool){
	if root == nil {
		*res = false
		return
	}
	now += root.Val
	flag := root.Left == nil && root.Right == nil
	if flag && now == target {
		*res = true
		return
	}
	if root.Left != nil {
		hasPathAux(root.Left, now, target, res)
	}
	if root.Right != nil {
		hasPathAux(root.Right, now, target, res)
	}
	now -= root.Val
}

func hasPathSum(root *types.TreeNode, sum int) bool {
	res := false
	hasPathAux(root, 0, sum, &res)
	return res
}
