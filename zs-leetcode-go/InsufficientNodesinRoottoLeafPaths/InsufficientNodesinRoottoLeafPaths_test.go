package InsufficientNodesinRoottoLeafPaths

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestSufficientSubset(t *testing.T) {
	tree := &types.TreeNode{
		Val:   1,
		Left:  &types.TreeNode{
			Val:   2,
			Left:  &types.TreeNode{
				Val:   -5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &types.TreeNode{
			Val:   -3,
			Left:  &types.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	root := sufficientSubset(tree, -1)
	root.Floor()
}