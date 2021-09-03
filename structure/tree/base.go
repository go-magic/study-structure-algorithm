package tree

import "fmt"

type Base struct {
	n INode
}

func NewBase(n INode) *Base {
	return &Base{
		n: n,
	}
}

/*
思考?为什么是用函数获取长度而不是对象的函数?
*/
func (b Base) Length() int {
	return b.length(b.n)
}

func (b Base) length(n INode) int {
	if b.n.Left() == nil && b.n.Right() == nil {
		return 0
	}
	leftLength := 0
	rightLength := 0
	if b.n.Left() != nil {
		leftLength = b.length(n.Left()) + 1
	}
	if b.n.Right() != nil {
		rightLength = b.length(n.Right()) + 1
	}
	if leftLength > rightLength {
		return leftLength
	}
	return rightLength
}

/*
查找最小节点
*/
func (b Base) FindMinNode() INode {
	return b.findMinNode(b.n.Left())
}

func (b Base) findMinNode(n INode) INode {
	if b.n.Left() == nil {
		return b.n.Left()
	}
	return b.findMinNode(n.Left())
}

/*
前序遍历
*/
func (b Base) PrintPre() {
	b.printPre(b.n)
}

func (b Base) printPre(n INode) {
	fmt.Println(n.GetData())
	if n.Left() != nil {
		b.printPre(n.Left())
	}
	if n.Right() != nil {
		b.printPre(n.Right())
	}
}

/*
中序遍历
*/
func (b Base) PrintMid() {
	b.printMid(b.n)
}

func (b Base) printMid(n INode) {
	if n.Left() != nil {
		b.printMid(n.Left())
	}
	fmt.Println(n.GetData())
	if n.Right() != nil {
		b.printMid(n.Right())
	}
}

/*
中序遍历
*/
func (b Base) PrintNext() {
	b.printNext(b.n)
}

func (b Base) printNext(n INode) {
	if n.Left() != nil {
		b.printNext(n.Left())
	}
	if n.Right() != nil {
		b.printNext(n.Right())
	}
	fmt.Println(n.GetData())
}
