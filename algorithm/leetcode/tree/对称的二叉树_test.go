package tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	n := NewTreeNode(1)
	n.Left = NewTreeNode(2)
	n.Right = NewTreeNode(2)
	n.Left.Left = NewTreeNode(3)
	n.Left.Right = NewTreeNode(4)
	n.Right.Left = NewTreeNode(4)
	n.Right.Right = NewTreeNode(3)
	if ok := isSymmetric(n); !ok {
		t.Fatal("测试不通过")
		return
	}
	n = NewTreeNode(1)
	n.Left = NewTreeNode(2)
	n.Right = NewTreeNode(2)
	n.Left.Right = NewTreeNode(3)
	n.Right.Right = NewTreeNode(3)
	if ok := isSymmetric(n); ok {
		t.Fatal("测试不通过")
		return
	}
	n = NewTreeNode(1)
	n.Left = NewTreeNode(2)
	n.Right = NewTreeNode(3)
	if ok := isSymmetric(n); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return equal(root.Left, root.Right)
}

func equal(left, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	if ok := equal(left.Left, right.Right); !ok {
		return false
	}
	return equal(left.Right, right.Left)
}
