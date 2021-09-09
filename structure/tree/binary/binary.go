package binary

import "github.com/go-magic/study-structure-algorithm/structure/tree/node"

type Node struct {
	left  *Node
	right *Node
	iData node.IData
}

func NewNode(data node.IData) *Node {
	return &Node{
		iData: data,
	}
}

func (n *Node) Left() node.INode {
	if n.left == nil {
		return nil
	}
	return n.left
}

func (n *Node) Right() node.INode {
	if n.right == nil {
		return nil
	}
	return n.right
}

func (n *Node) Data() node.IData {
	return n.iData
}

func (n *Node) Insert(data node.IData) {
	if data < n.iData {
		if n.left == nil {
			n.left = NewNode(data)
			return
		}
		n.left.Insert(data)
	}
	if data > n.iData {
		if n.right == nil {
			n.right = NewNode(data)
			return
		}
		n.right.Insert(data)
	}
}

func (n *Node) Delete(data node.IData) node.INode {
	return n.delete(data)
}

func (n *Node) delete(data node.IData) node.INode {
	dNode := node.FindNode(n, data)
	if dNode.Right() == nil {
		return dNode.Left()
	}
	rNode := node.FindMinNode(dNode.Right())
	rNode.Left() =
}

func (n *Node)AddLeft(left node.INode)  {

}

func (n *Node)AddRight(right node.INode)  {

}
