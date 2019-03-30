package SameTree

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	root1 := &types.TreeNode{1, &types.TreeNode{2, nil, nil}, nil}
	root2 := &types.TreeNode{1, nil, &types.TreeNode{2, nil, nil}}
	fmt.Println(isSameTree(root1, root2))
	root1 = &types.TreeNode{1, &types.TreeNode{2, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(isSameTree(root1, root1))
}
