package BinarySearchTreeIterator

import (
    "testing"
	"fmt"

    "for-an-offer/zs-leetcode-go/types"
)
func TestIterator(t *testing.T) {
    tree := &types.TreeNode{
        Val: 7,
        Left: &types.TreeNode{
            Val: 3, 
            Left: nil,
            Right: nil,
        },
        Right: &types.TreeNode{
            Val: 15,
            Left: &types.TreeNode{
                Val: 9, 
                Left: nil,
                Right: nil,
            },
            Right: &types.TreeNode{
                Val: 20,
                Left: nil,
                Right: nil,
            },
        },
    }
    obj := Constructor(tree)
    fmt.Println(obj.Next())
    fmt.Println(obj.Next())
    fmt.Println(obj.HasNext())
    fmt.Println(obj.Next())
    fmt.Println(obj.HasNext())
    fmt.Println(obj.Next())
    fmt.Println(obj.HasNext())
    fmt.Println(obj.Next())
    fmt.Println(obj.HasNext())

}
