package SumOfLeftLeaves

import (
    "fmt"
    "testing"
    "for-an-offer/zs-leetcode-go/types"
)

func TestSumOfLeftleaves(t *testing.T) {
    tree := &types.TreeNode{
        Left: &types.TreeNode{
            Left: nil,
            Right: nil,
            Val: 9,
        },
        Right: &types.TreeNode{
            Left: &types.TreeNode{
                Left: nil,
                Right: nil,
                Val: 15,
            },
            Right: &types.TreeNode{
                Left: nil,
                Right: nil,
                Val: 7,
            },
            Val: 20,
        },
        Val: 3,
    }
    fmt.Println(sumOfLeftLeaves(tree))
    tree1 := &types.TreeNode{
        Left: nil,
        Right: nil,
        Val: 2,
    }
    fmt.Println(sumOfLeftLeaves(tree1))
    tree2 := &types.TreeNode{
        Left: &types.TreeNode{
            Left: nil,
            Right: nil,
            Val: 8,
        },
        Right: nil,
        Val: 8,
    }
    fmt.Println(sumOfLeftLeaves(tree2))
}
