package node

type IData int

type INode interface {
	Left() INode
	AddLeft(INode)
	Right() INode
	AddRight(INode)
	Data() IData
	Insert(IData)
	Delete(IData) INode
}

func FindMinNode(n INode) INode {
	if n.Left() == nil {
		return n
	}
	return FindMinNode(n.Left())
}

func FindNodeLength(n INode) int {
	var l, r int
	if n.Left() != nil {
		l = FindNodeLength(n.Left()) + 1
	}
	if n.Right() != nil {
		r = FindNodeLength(n.Right()) + 1
	}
	if l > r {
		return l
	}
	return r
}

func FindNode(n INode, data IData) INode {
	if n.Data() == data {
		return n
	}
	if data < n.Data() {
		return FindNode(n.Left(), data)
	}
	if data > n.Data() {
		return FindNode(n.Right(), data)
	}
	return nil
}
