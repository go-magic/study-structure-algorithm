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

func (n *Node) delete(data node.IData) *Node {
	dNode := n.FindNode(data)
	if dNode.left == nil {
		return dNode.left
	}
	rNode := n.FindMinNode()
	rNode.left = dNode.left
	return rNode
}

func (n *Node) FindNode(data node.IData) *Node {
	if n.Data() == data {
		return n
	}
	if data < n.Data() {
		return n.left.FindNode(data)
	}
	if data > n.Data() {
		return n.right.FindNode(data)
	}
	return nil
}

func (n *Node) FindMinNode() *Node {
	if n.left == nil {
		return n
	}
	return n.left.FindMinNode()
}

func (n *Node) AddLeft(left node.INode) {

}

func (n *Node) AddRight(right node.INode) {

}
