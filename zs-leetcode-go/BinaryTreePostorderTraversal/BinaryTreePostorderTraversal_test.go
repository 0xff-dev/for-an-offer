package BinaryTreePostorderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := &types.TreeNode{1, nil, &types.TreeNode{2,
		&types.TreeNode{3, nil, nil}, nil}}
	fmt.Println(postorderTraversal(root))
}
