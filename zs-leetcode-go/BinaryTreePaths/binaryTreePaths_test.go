package BinaryTreePaths

import (
    "fmt"
    "testing"

    "for-an-offer/zs-leetcode-go/types"
)

func TestBinaryTreePaths(t *testing.T) {
    tree := &types.TreeNode{
        Right: &types.TreeNode{
            Left: nil,
            Right: nil,
            Val: 3,
        },
        Left: &types.TreeNode{
            Left: nil,
            Right: &types.TreeNode{
                Left: nil,
                Right: nil,
                Val: 5,
            },
            Val: 2,
        },
        Val: 1,
    }
    path := binaryTreePaths(tree)
    fmt.Println(path)
    tree = &types.TreeNode{
        Left: nil,
        Right: nil,
        Val: 1,
    }
    fmt.Println(binaryTreePaths(tree))
}
