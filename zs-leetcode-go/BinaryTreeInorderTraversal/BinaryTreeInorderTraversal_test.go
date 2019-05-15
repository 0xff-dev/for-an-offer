package BinaryTreeInorderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &types.TreeNode{1, nil,
		&types.TreeNode{2, &types.TreeNode{3, nil, nil}, nil}}
	fmt.Println(inorderTraversal(root))
}
