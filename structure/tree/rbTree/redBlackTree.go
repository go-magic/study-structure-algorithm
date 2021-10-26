package rbTree

type RedBlackTree struct {
	root *RedBlackNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (r *RedBlackTree) Insert(data NodeData) {
	if r.root != nil {
		r.root = r.root.Insert(data)
	} else {
		r.root = NewRedBlackNode(data)
	}
	r.root.color = Black
}
