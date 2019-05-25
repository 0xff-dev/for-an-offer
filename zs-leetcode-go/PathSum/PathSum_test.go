package PathSum

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tree := &types.TreeNode{5,
		&types.TreeNode{4, &types.TreeNode{11,
			&types.TreeNode{7, nil, nil},
			&types.TreeNode{2, nil, nil}}, nil},
		&types.TreeNode{8,
			&types.TreeNode{13, nil, nil},
			&types.TreeNode{4, nil, &types.TreeNode{1, nil, nil}}}}
	fmt.Println(hasPathSum(tree, 22))
}
