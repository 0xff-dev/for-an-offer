package PathSumII

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestPathSum(t *testing.T) {
	//tree := &types.TreeNode{5,
	//	&types.TreeNode{4, &types.TreeNode{11,
	//		&types.TreeNode{7, nil, nil},
	//		&types.TreeNode{2, nil, nil}}, nil},
	//	&types.TreeNode{8,
	//		&types.TreeNode{13, nil, nil},
	//		&types.TreeNode{4, nil, &types.TreeNode{1, nil, nil}}}}
	//fmt.Println(pathSum(tree, 22))
	tree := &types.TreeNode{5, &types.TreeNode{4, &types.TreeNode{11,
			&types.TreeNode{7, nil, nil},
			&types.TreeNode{2, nil, nil}}, nil},
		&types.TreeNode{8,
			&types.TreeNode{13, nil, nil},
			&types.TreeNode{4, &types.TreeNode{5, nil, nil},
				&types.TreeNode{1, nil, nil}}}}
	fmt.Println(pathSum(tree, 22))

	tree = &types.TreeNode{-2, nil, &types.TreeNode{-3, nil, nil}}
	fmt.Println(pathSum(tree, -5))
	fmt.Println(pathSum(nil, 2))

	tree = &types.TreeNode{1, &types.TreeNode{-2, nil, nil},
		&types.TreeNode{3, nil, nil}}
	fmt.Println(pathSum(tree, -1))

	tree = &types.TreeNode{1, &types.TreeNode{-2,
		&types.TreeNode{1, &types.TreeNode{-1, nil, nil}, nil},
			&types.TreeNode{3, nil, nil}}, &types.TreeNode{-3, &types.TreeNode{-2, nil, nil}, nil}}
	fmt.Println(pathSum(tree, -1))

	tree = &types.TreeNode{1, &types.TreeNode{0, nil, nil},
		&types.TreeNode{1, nil, nil}}
	tree.Left.Left = &types.TreeNode{1, &types.TreeNode{0, nil, nil},
		&types.TreeNode{1, nil, nil}}
	tree.Left.Right = &types.TreeNode{2, &types.TreeNode{-1, nil, nil},
		&types.TreeNode{0, nil, nil},  }

	tree.Right.Left = &types.TreeNode{0, &types.TreeNode{-1, nil, nil},
		&types.TreeNode{0, nil, nil}, }
	tree.Right.Right = &types.TreeNode{-1, &types.TreeNode{1, nil, nil},
		&types.TreeNode{0, nil, nil}, }

	fmt.Println(pathSum(tree, 2))
}
