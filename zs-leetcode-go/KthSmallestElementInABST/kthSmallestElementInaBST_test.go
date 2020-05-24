package KthSmallestElementInABST

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tree := &types.TreeNode{
		Val:   3,
		Left:  &types.TreeNode{
			Val:   1,
			Left:  nil,
			Right: &types.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &types.TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(tree, 1))
	tree = &types.TreeNode{
		Val:   5,
		Left:  &types.TreeNode{
			Val:   3,
			Left:  &types.TreeNode{
				Val:   2,
				Left:  &types.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &types.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &types.TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(tree, 3))
}
