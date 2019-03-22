package FullBinaryTree

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestIsFullBinaryTree(t *testing.T) {
	tree := &types.TreeNode{1, nil, nil}
	fmt.Println(IsFullBinaryTree(tree))
	tree = &types.TreeNode{4, &types.TreeNode{4, nil, nil}, nil}
	fmt.Println(IsFullBinaryTree(tree))
	tree = &types.TreeNode{5, &types.TreeNode{4, nil, nil},
	&types.TreeNode{6, nil, nil}}
	fmt.Println(IsFullBinaryTree(tree))
}
