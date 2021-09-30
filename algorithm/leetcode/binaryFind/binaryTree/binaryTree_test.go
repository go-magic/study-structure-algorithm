package binarytree

import "testing"

func TestBinaryTree(t *testing.T) {
	tree := NewTreeNode(1)
	tree.Right = NewTreeNode(2)
	tree.Right.Left = NewTreeNode(3)
	arr := inorderTraversal(tree)
	t.Log(arr)
}
