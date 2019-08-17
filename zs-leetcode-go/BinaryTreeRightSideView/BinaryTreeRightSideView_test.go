package BinaryTreeRightSideView

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := &types.TreeNode{1,
		&types.TreeNode{2, nil, &types.TreeNode{5, nil, nil}},
		&types.TreeNode{3, nil, &types.TreeNode{4, nil, nil}}}
	fmt.Println(rightSideView(root))
}
