package MaximumDepthofBinaryTree

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &types.TreeNode{3, nil, nil}
	root.Left = &types.TreeNode{9, nil, nil}
	root.Right = &types.TreeNode{20, &types.TreeNode{15, nil, nil},
		&types.TreeNode{7, nil, nil}}
	fmt.Println(maxDepth(root))
}
