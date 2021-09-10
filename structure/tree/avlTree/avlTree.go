package avl

import "github.com/go-magic/study-structure-algorithm/structure/tree/node"

type Tree struct {
	root *node.Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = node.NewNode(data)
		return
	}
	t.root.Insert(data)
}
