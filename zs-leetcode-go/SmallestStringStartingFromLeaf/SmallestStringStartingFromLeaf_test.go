package SmallestStringStartingFromLeaf

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestSmallestFromLeaf(t *testing.T) {
	root := &types.TreeNode{Val: 0,
		Left: &types.TreeNode{Val: 1,
			Left:  &types.TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &types.TreeNode{Val: 4, Left: nil, Right: nil}},
		Right: &types.TreeNode{Val: 2,
			Left:  &types.TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &types.TreeNode{Val: 4, Left: nil, Right: nil}}}
	fmt.Println(smallestFromLeaf(root))
	root = &types.TreeNode{Val: 25,
		Left: &types.TreeNode{Val: 1,
			Left:  &types.TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &types.TreeNode{Val: 3, Left: nil, Right: nil}},
		Right: &types.TreeNode{Val: 3,
			Left:  &types.TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &types.TreeNode{Val: 2, Left: nil, Right: nil}}}
	fmt.Println(smallestFromLeaf(root))
}
