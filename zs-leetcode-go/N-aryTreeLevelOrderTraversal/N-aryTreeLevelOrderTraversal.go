package N_aryTreeLevelOrderTraversal

import "for-an-offer/zs-leetcode-go/types"

func levelOrder(root *types.Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*types.Node{root}
	for len(queue) > 0 {
		tmp := make([]*types.Node, 0)
		tmpArr := make([]int, 0)
		for _, item := range queue {
			tmpArr = append(tmpArr, item.Val)
			if len(item.Children) > 0 {
				tmp = append(tmp, item.Children...)
			}
		}
		res = append(res, tmpArr)
		queue = tmp
	}
	return res
}
