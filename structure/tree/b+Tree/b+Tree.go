package b_Tree

type Tree struct {
	root *Block
}

func NewTree(m int) *Tree {
	M = m
	return &Tree{}
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = NewBlock(data)
		return
	}
	t.root.Insert(data)
	if t.root.Length == M {
		node := t.root.getMidNode()
		block := NewBlock(node.Data)
		block.Head.LeftChild = t.root
		block.Head.LeftChild.Length = M / 2
		block.Head.RightChild = &Block{Head: node}
		block.Head.RightChild.Length = M - M/2
		t.root = block
	}
}
