package binary

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

func (t *Tree) FindMinData() int {
	minNode := t.root.FindMinNode()
	return minNode.Data
}

func (t *Tree) Length() int {
	return t.root.Length()
}

func (t *Tree) Delete(data int) {
	t.root = t.root.Delete(data)
}

func (t *Tree) NodeNum() int {
	if t.root == nil {
		return 0
	}
	return t.root.NodeNum() + 1
}
