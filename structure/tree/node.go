package tree

type INode interface {
	GetData() int
	Left() INode
	Right() INode
	Insert(int)
	Delete(int) INode
}

/*
思考?为什么是用函数获取长度而不是对象的函数?
*/
func Length(n INode) int {
	if n.Left() == nil && n.Right() == nil {
		return 0
	}
	leftLength := 0
	rightLength := 0
	if n.Left() != nil {
		leftLength = Length(n.Left()) + 1
	}
	if n.Right() != nil {
		rightLength = Length(n.Right()) + 1
	}
	if leftLength > rightLength {
		return leftLength
	}
	return rightLength
}
