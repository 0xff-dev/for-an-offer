package ConstructBinaryTreefromInorderandPostorderTraversal

import "testing"

func TestBuildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tree := buildTree(inorder, postorder)
	tree.Floor()
}
