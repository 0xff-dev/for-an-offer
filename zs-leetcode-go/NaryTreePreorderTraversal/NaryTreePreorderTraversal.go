package NaryTreePreorderTraversal

import "for-an-offer/zs-leetcode-go/types"

func preorder(root *types.Node) []int {
	res := make([]int, 0)
	aux(root, &res)
	return res
}

func aux(root *types.Node, array *[]int) {
	if root == nil {
		return
	}
	*array = append(*array, root.Val)
	for _, children := range root.Children {
		aux(children, array)
	}
}
