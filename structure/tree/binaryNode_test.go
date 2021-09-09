package tree

import (
	"github.com/go-magic/study-structure-algorithm/structure/tree/binary"
	"github.com/go-magic/study-structure-algorithm/structure/tree/node"
	"testing"
)

func createNode() node.INode {
	n := binary.NewNode(10)
	n.Insert(1)
	n.Insert(3)
	n.Insert(5)
	n.Insert(19)
	n.Insert(4)
	n.Insert(21)
	n.Insert(6)
	n.Insert(7)
	return n
}

/*
按行打印
*/
func TestPrint(t *testing.T) {
	t.Log(node.FindNodeLength(createNode()))
}

/*
测试层高
*/
func TestFindNode(t *testing.T) {

}

/*
测试层高
*/
func TestLength(t *testing.T) {

}

/*
踩坑
*/
func newINode() INode {
	return (*BinaryNode)(nil)
}

func TestNil(t *testing.T) {
	i := newINode()
	if i == nil {
		t.Log("nil")
		return
	}
	t.Log("not nil")
}
