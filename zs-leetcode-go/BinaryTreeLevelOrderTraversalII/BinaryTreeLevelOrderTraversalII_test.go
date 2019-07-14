package BinaryTreeLevelOrderTraversalII

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	root := &types.TreeNode{3,
		&types.TreeNode{9, nil, nil},
	&types.TreeNode{20,
		&types.TreeNode{15, nil, nil},
	&types.TreeNode{7, nil, nil}}}
	fmt.Println(levelOrderBottom(root))
}
