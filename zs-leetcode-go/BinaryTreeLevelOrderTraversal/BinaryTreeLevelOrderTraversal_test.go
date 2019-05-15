package BinaryTreeLevelOrderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := &types.TreeNode{Val: 3, Left: &types.TreeNode{Val: 9, Left: nil, Right: nil}, Right: nil}
	tree.Right = &types.TreeNode{Val: 20, Left: &types.TreeNode{Val: 15, Left: nil, Right: nil},
		Right: &types.TreeNode{Val: 7, Left: nil, Right: nil}}
	fmt.Println(levelOrder(tree))
}
