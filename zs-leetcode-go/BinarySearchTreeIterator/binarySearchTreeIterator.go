package BinarySearchTreeIterator

import (
    "fmt"
    "for-an-offer/zs-leetcode-go/types"
)


// 一个简单的办法就是 构造的时候直接前序遍历，然后放入队列里面。后面直接利用队列就行，就是
// 就是这个方法
type BSTIterator struct {    
    root *types.TreeNode
    array []*types.TreeNode
}


func (self *BSTIterator) inorder(root *types.TreeNode) {
    if root.Left != nil {
        self.inorder(root.Left)
    }
    self.array = append(self.array, root)
    if root.Right != nil {
        self.inorder(root.Right)
    }
}

func Constructor(root *types.TreeNode) BSTIterator {
    obj := BSTIterator{
        root: root,
        array: make([]*types.TreeNode, 0),
    }
    if root != nil {
        (&obj).inorder(root)
    }
    return obj
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    top := this.array[0]
    this.array = this.array[1:]
    return top.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
   return len(this.array) > 0 
}


//二叉搜索树，前序遍历是升序
