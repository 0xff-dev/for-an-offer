package N_aryTreeLevelOrderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := &types.Node{
		Val:      1,
		Children: []*types.Node{
			{
				Val: 3,
				Children: []*types.Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
				Children: nil,
			},
			{
				Val: 4,
				Children: nil,
			},
		},
	}
	fmt.Println(levelOrder(tree))
}
