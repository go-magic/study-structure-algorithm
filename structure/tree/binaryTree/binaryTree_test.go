package binary

import "testing"

func initBinaryTree() *Tree {
	tree := NewTree()
	tree.Insert(200)
	tree.Insert(160)
	tree.Insert(240)
	tree.Insert(120)
	tree.Insert(180)
	tree.Insert(220)
	tree.Insert(280)
	tree.Insert(80)
	tree.Insert(40)
	tree.Insert(100)
	tree.Insert(90)
	return tree
}

func TestNewTree(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(80)
	tree.Delete(90)
}
