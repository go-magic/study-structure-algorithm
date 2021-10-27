package tree

import "testing"

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	image := NewTreeNode(root.Val)
	image.Right = mirrorTree(root.Left)
	image.Left = mirrorTree(root.Right)
	return image
}

func TestMirrorTree(t *testing.T) {
	n := NewTreeNode(4)
	n.Left = NewTreeNode(2)
	n.Right = NewTreeNode(7)
	n.Left.Left = NewTreeNode(1)
	n.Left.Right = NewTreeNode(3)
	n.Right.Left = NewTreeNode(6)
	n.Right.Right = NewTreeNode(9)
	m := mirrorTree(n)

	if m.Val != 4 {
		t.Fatal("测试失败")
		return
	}
	if m.Left.Val != 7 {
		t.Fatal("测试失败")
		return
	}
	if m.Right.Val != 2 {
		t.Fatal("测试失败")
		return
	}
	if m.Left.Left.Val != 9 {
		t.Fatal("测试失败")
		return
	}
	if m.Left.Right.Val != 6 {
		t.Fatal("测试失败")
		return
	}
	if m.Right.Left.Val != 3 {
		t.Fatal("测试失败")
		return
	}
	if m.Right.Right.Val != 1 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
