package MinimumDepthofBinaryTree111

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func  TestMinimumDepthofBinaryTree(t *testing.T) {
	root := &types.TreeNode{
		Val: 1,
		Left: &types.TreeNode{Val: 2, Left: nil, Right: nil},
		Right: nil,
	}
	fmt.Println(MinimumDepthofBinaryTree(root))
}
