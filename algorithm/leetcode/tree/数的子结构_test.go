package tree

import "testing"

func TestiIsSubStructure(t *testing.T) {
	a := initTree()
	b := initTree()
	trees := isSubStructure(a, b)
	t.Log("测试通过", trees)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return false
}

func findNode(val int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	var n *TreeNode
	if node.Left != nil {
		n = findNode(val, node.Left)
	}
	if n != nil {
		return n
	}
	if node.Right != nil {
		n = findNode(val, node.Right)
	}
	return n
}

func TestFindNode(t *testing.T) {
	n := NewTreeNode(10)
	n.Left = NewTreeNode(8)
	n.Right = NewTreeNode(14)
	n.Left.Left = NewTreeNode(4)
	n.Left.Right = NewTreeNode(9)
	n.Right.Left = NewTreeNode(18)
	n.Right.Right = NewTreeNode(24)
	r := findNode(14, n)
	t.Log(r)
}
