package tree

import "testing"

func TestiIsSubStructure(t *testing.T) {
	a := initTree()
	b := initTree()
	trees := isSubStructure(a, b)
	t.Log("测试通过", trees)
}

var ret = make([]*TreeNode, 0)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	ret = make([]*TreeNode, 0)
	findNode(B.Val, A)
	for _, node := range ret {
		if ok := exist(node, B); ok {
			return true
		}
	}
	return false
}

func exist(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	if ok := exist(A.Left, B.Left); !ok {
		return false
	}
	return exist(A.Right, B.Right)
}

func findNode(val int, A *TreeNode) {
	if A == nil {
		return
	}
	if A.Val == val {
		ret = append(ret, A)
	}
	findNode(val, A.Left)
	findNode(val, A.Right)
}

func TestFindNode(t *testing.T) {
	n := NewTreeNode(3)
	n.Left = NewTreeNode(4)
	n.Right = NewTreeNode(5)
	n.Left.Left = NewTreeNode(4)
	n.Left.Right = NewTreeNode(2)
	n.Left.Left.Left = NewTreeNode(1)

	b := NewTreeNode(4)
	b.Left = NewTreeNode(1)
	if ok := isSubStructure(n, b); !ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}
