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
	createNode().printRows()
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
