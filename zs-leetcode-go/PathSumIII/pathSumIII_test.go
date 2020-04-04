package PathSumIII

import (
    "fmt"
    "testing"

    "for-an-offer/zs-leetcode-go/types"
)

func TestPathSum(t *testing.T) {
    tree := &types.TreeNode{
        Left: &types.TreeNode{
            Left: &types.TreeNode{
                Left: &types.TreeNode{
                    Left: nil,
                    Right: nil,
                    Val: 3,
                },
                Right: &types.TreeNode{
                    Left: nil,
                    Right: nil,
                    Val: -2,
                },
                Val: 3,
            },
            Right: &types.TreeNode{
                Left: nil, 
                Right: &types.TreeNode{
                    Left: nil,
                    Right: nil,
                    Val: 1,
                },
                Val: 2,
            },
            Val: 5,
        },
        Right: &types.TreeNode{
            Left: nil,
            Right: &types.TreeNode{
                Left: nil,
                Right: nil,
                Val: 11,
            },
            Val: -3,
        },
        Val: 10,
    }
    fmt.Println(pathSum(tree, 8))
    fmt.Println(pathSum(nil, 8))
}

