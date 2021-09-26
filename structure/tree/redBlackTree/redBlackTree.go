package redBlackTree

import "github.com/go-magic/study-structure-algorithm/structure/tree/node"

type Tree struct {
	root *node.RedBlackNode
}

func NewRedBlackTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = node.NewRedBlackNode(data)
		t.root.SetBlackColor()
		return
	}
	t.root = t.root.Insert(data)
	t.root.SetBlackColor()
}

func (t *Tree) Length() int {
	if t.root == nil {
		return 0
	}
	return t.root.Length() + 1
}

func (t *Tree) NodeNum() int {
	if t.root == nil {
		return 0
	}
	return t.root.NodeNum() + 1
}
