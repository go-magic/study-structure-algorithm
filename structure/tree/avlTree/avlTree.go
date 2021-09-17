package avl

import "github.com/go-magic/study-structure-algorithm/structure/tree/node"

type Tree struct {
	root *node.AvlNode
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = node.NewAvlNode(data)
		return
	}
	t.root = t.root.Insert(data)
}

func (t *Tree) Length() int {
	if t.root == nil {
		return 0
	}
	return t.root.Length()
}

func (t *Tree) NodeNum() int {
	if t.root == nil {
		return 0
	}
	return t.root.NodeNum() + 1
}

func (t *Tree) Delete(data int) {
	t.root = t.root.DeleteNode(data)
}

func (t *Tree) RecursionLength() int {
	if t.root == nil {
		return 0
	}
	return t.root.RecursionLength() + 1
}
