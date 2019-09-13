package SumRoottoLeafNumbers

import (
	"fmt"
	"for-an-offer/zs-leetcode-go/types"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	root := &types.TreeNode{Val: 1, Left: &types.TreeNode{Val: 2, Left: nil, Right: nil}, Right: &types.TreeNode{Val: 3, Left: nil, Right: nil}}
	fmt.Println(sumNumbers(root))
	root = &types.TreeNode{Val: 4,
		Left:  &types.TreeNode{Val: 9, Left: &types.TreeNode{Val: 5, Left: nil, Right: nil}, Right: &types.TreeNode{Val: 1, Left: nil, Right: nil}},
		Right: &types.TreeNode{Val: 0, Left: nil, Right: nil}}
	fmt.Println(sumNumbers(root))
	fmt.Println(sumNumbers(&types.TreeNode{Val: 10, Left: nil, Right: nil}))
	fmt.Println(sumNumbers(nil))
}
