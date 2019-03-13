package InvertBinaryTree226

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	root := &types.TreeNode{
		Val:10,
		Left: &types.TreeNode{
			Val: 6,
			Left: &types.TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
			Right: &types.TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
		},
		Right: &types.TreeNode{
			Val: 7,
			Left: nil,
			Right: nil,
		},
	}
	root.Floor()
	InvertBinaryTree(root)
	root.Floor()
}
