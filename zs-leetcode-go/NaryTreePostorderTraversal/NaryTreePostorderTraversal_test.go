package NaryTreePostorderTraversal

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestPostOrder(t *testing.T) {
	tree := &types.Node{
		Val:      1,
		Children: []*types.Node{
			{
				3,
				[]*types.Node{
					{
						5,
						nil,
					},
					{
						6,
						nil,
					},
				},
			},
			{
				2,
				nil,
			},
			{
				4,
				nil,
			},
		},
	}
	fmt.Println(postorder(tree))
}
