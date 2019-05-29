package BinaryTreeZigzagLevelOrderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tree := &types.TreeNode{2, &types.TreeNode{9, nil, nil},
		&types.TreeNode{20, &types.TreeNode{15, nil, nil},
			&types.TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(tree))

	tree = &types.TreeNode{1, &types.TreeNode{2, &types.TreeNode{4, nil, nil} ,nil},
		&types.TreeNode{3, nil, &types.TreeNode{5, nil, nil}}}
	fmt.Println(zigzagLevelOrder(tree))
}
