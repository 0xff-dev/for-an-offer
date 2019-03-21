package examples

import (
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestConvertBinaryToDoubleList(t *testing.T) {
	tree := &types.TreeNode{2, &types.TreeNode{3, nil, nil}, &types.TreeNode{
		4, &types.TreeNode{1, nil, nil}, nil}}
	ConvertBinaryToDoubleList(tree)
}
