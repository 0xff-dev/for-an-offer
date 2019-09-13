package SmallestStringStartingFromLeaf

import (
	"bytes"
	"for-an-offer/zs-leetcode-go/types"
)

func min(obj1 []int, obj2 string) string {
	buf := bytes.NewBufferString("")
	for index := len(obj1) - 1; index >= 0; index-- {
		buf.WriteByte(byte(obj1[index] + 97))
	}
	str := buf.String()
	if str > obj2 {
		return obj2
	}
	return str
}

func aux(node *types.TreeNode, now []int, res *string) {
	flag := false
	if node.Left != nil {
		aux(node.Left, append(now, node.Left.Val), res)
		flag = true
	}
	if node.Right != nil {
		aux(node.Right, append(now, node.Right.Val), res)
		flag = true
	}
	if !flag {
		*res = min(now, *res)
	}
}
func smallestFromLeaf(root *types.TreeNode) string {
	res := "z"
	now := make([]int, 0)
	aux(root, append(now, root.Val), &res)
	return res
}
