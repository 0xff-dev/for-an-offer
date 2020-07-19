package MinimumDistanceBetweenBSTNodes

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestMinDiffInBST(t *testing.T) {
	tree := &types.TreeNode{
		Val:   4,
		Left:  &types.TreeNode{
			Val:   2,
			Left:  &types.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &types.TreeNode{
				Val:   3,
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
	fmt.Println(minDiffInBST(tree))
}
