package tree

import (
	"testing"
)

func createNode() *BinaryNode {
	node := newBinaryNode(10)
	node.Insert(1)
	node.Insert(3)
	node.Insert(5)
	node.Insert(19)
	node.Insert(4)
	node.Insert(21)
	node.Insert(6)
	node.Insert(7)
	return node
}

/*
按行打印
*/
func TestPrint(t *testing.T) {
	tree := NewTree(createNode(), rowPrint{})
	t.Log(tree.FindMinNode())
	tree.PrintRows()
}

/*
测试层高
*/
func TestFindNode(t *testing.T) {
	node := createNode()
	n := node.findNode(21)
	t.Log(n.data)
}

/*
测试层高
*/
func TestLength(t *testing.T) {
	node := createNode()
	t.Log(Length(node))
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
