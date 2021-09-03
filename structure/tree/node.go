package tree

type INode interface {
	GetData() int
	Left() INode
	Right() INode
}
