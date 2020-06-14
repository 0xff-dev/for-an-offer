package NaryTreePostorderTraversal

import "for-an-offer/zs-leetcode-go/types"

func postorder(root *types.Node) []int {
	res := make([]int, 0)
	aux(root, &res)
	return res
}


func aux(root *types.Node, array *[]int) {
	if root == nil {
		return
	}
	for _, children := range root.Children {
		aux(children, array)
	}
	*array = append(*array, root.Val)
}