package FlattenBinaryTreetoLinkedList

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestFlatten(t *testing.T) {
	tree := &types.TreeNode{1,
		&types.TreeNode{2,
			&types.TreeNode{3, nil, nil}, &types.TreeNode{4, nil, nil}},
		&types.TreeNode{5, nil, &types.TreeNode{6, nil, nil}}}
	fmt.Println("test")
	flatten(tree)
	for ; tree != nil; tree = tree.Right {
		fmt.Println(tree.Val)
	}
	tree = &types.TreeNode{1, &types.TreeNode{2, nil, nil},
		&types.TreeNode{5, nil, nil}}
	flatten(tree)
	for ; tree != nil; tree = tree.Right {
		fmt.Println(tree.Val)
	}
}
