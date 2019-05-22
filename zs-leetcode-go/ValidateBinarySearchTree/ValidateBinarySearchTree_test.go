package ValidateBinarySearchTree

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	root := &types.TreeNode{2, &types.TreeNode{1, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(isValidBST(root))
	root = &types.TreeNode{5, &types.TreeNode{1, nil, nil},
		&types.TreeNode{4, &types.TreeNode{3, nil, nil},
			&types.TreeNode{6, nil, nil}}}
	fmt.Println(isValidBST(root))
	root = &types.TreeNode{1, &types.TreeNode{1, nil, nil}, nil}
	fmt.Println(isValidBST(root))
}
