package MinimumDistanceBetweenBSTNodes

import "for-an-offer/zs-leetcode-go/types"

func minDiffInBST(root *types.TreeNode) int {
	if root == nil {
		return 0
	}
	var node *types.TreeNode
	aux(root, &node)
	min := 1000000
	walker := node
	for walker.Left != nil {
		if walker.Val - walker.Left.Val < min {
			min = walker.Val - walker.Left.Val
		}
		walker = walker.Left
	}
	return min
}

func aux(root *types.TreeNode, walker **types.TreeNode) {
	if root == nil {
		return
	}
	cur := root
	if cur.Left != nil {
		aux(cur.Left, walker)
	}
	cur.Left = *walker
	if *walker != nil {
		(*walker).Right = cur
	}
	*walker = cur
	if cur.Right != nil {
		aux(cur.Right, walker)
	}
}