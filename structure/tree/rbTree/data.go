package rbTree

type DefaultData int

func (d DefaultData) Hash() int {
	return int(d)
}

type NodeData interface {
	Hash() int
}
