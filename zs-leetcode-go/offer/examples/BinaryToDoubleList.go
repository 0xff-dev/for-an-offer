package examples

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
)

func ConvertBinaryToDoubleList(tree *types.TreeNode) {
	var lastnode *types.TreeNode
	auxFunc(tree, &lastnode)
	node := &lastnode
	for *node != nil {
		fmt.Println((*node).Val)
		*node = (*node).Left
	}
}

func auxFunc(tree *types.TreeNode, lastNode **types.TreeNode) {
	if tree == nil {
		return
	}
	pCurrent := tree
	if pCurrent.Left != nil {
		auxFunc(pCurrent.Left, lastNode)
	}
	pCurrent.Left = *lastNode
	if *lastNode != nil {
		(*lastNode).Right = pCurrent
	}
	*lastNode = pCurrent
	if pCurrent.Right != nil {
		auxFunc(pCurrent.Right, lastNode)
	}
}
