package redBlackTree

import (
	"github.com/go-magic/study-structure-algorithm/structure/tree/node"
	"testing"
)

func initBinaryTree() *Tree {
	tree := NewRedBlackTree()
	tree.Insert(200)
	tree.Insert(160)
	tree.Insert(240)
	tree.Insert(120)
	tree.Insert(180)
	tree.Insert(220)
	tree.Insert(280)
	tree.Insert(80)
	tree.Insert(130)
	tree.Insert(170)
	tree.Insert(40)
	return tree
}

func TestRule2(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(200)
	tree.Insert(160)
	if tree.root.Color == node.Black &&
		tree.root.Data == 200 &&
		tree.root.Left.Color == node.Red &&
		tree.root.Left.Data == 160 {
		t.Log("规则2通过")
		return
	}
	t.Fatal("规则2测试不通过")
}

func TestRule3(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(200)
	tree.Insert(160)
	tree.Insert(240)
	tree.Insert(120)
	if tree.root.Color == node.Black &&
		tree.root.Data == 200 &&
		tree.root.Left.Color == node.Red &&
		tree.root.Left.Data == 160 {
		t.Log("规则2通过")
		return
	}
	t.Fatal("规则2测试不通过")
}
