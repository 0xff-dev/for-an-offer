package BinaryTreeMaximumPathSum

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := &types.TreeNode{1, &types.TreeNode{2, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(maxPathSum(root))

	root = &types.TreeNode{-10, &types.TreeNode{9, nil, nil}, &types.TreeNode{
		20, &types.TreeNode{15, nil, nil},
		&types.TreeNode{7, nil, nil},
	}}
	fmt.Println(maxPathSum(root))
	root = &types.TreeNode{-3, nil, nil}
	fmt.Println(maxPathSum(root))
	root = &types.TreeNode{-1, &types.TreeNode{-2, nil, nil},
		&types.TreeNode{5, nil, nil}}
	fmt.Println(maxPathSum(root))
	root = &types.TreeNode{5,
		&types.TreeNode{4, &types.TreeNode{11,
			&types.TreeNode{7, nil, nil},
			&types.TreeNode{2, nil, nil}}, nil},
		&types.TreeNode{8, &types.TreeNode{13,
			nil, nil},
			&types.TreeNode{4, nil,
				&types.TreeNode{1, nil, nil}}}}
	fmt.Println(maxPathSum(root))
	root = &types.TreeNode{1, &types.TreeNode{-2, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(maxPathSum(root))
}
