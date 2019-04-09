package SymmetricTree

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := &types.TreeNode{1, nil, nil}
	root.Left = &types.TreeNode{2, &types.TreeNode{3, nil, nil},
		&types.TreeNode{4, nil, nil}}
	root.Right = &types.TreeNode{2, &types.TreeNode{4, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(isSymmetric(root))
	root = &types.TreeNode{1, nil, nil}
	root.Left = &types.TreeNode{2, nil, &types.TreeNode{3, nil, nil}}
	root.Right = &types.TreeNode{2, nil, &types.TreeNode{3, nil, nil}}
	fmt.Println(isSymmetric(root))
}
