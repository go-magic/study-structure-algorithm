package binary

import (
	"math/rand"
	"testing"
	"time"

	"github.com/go-magic/study-structure-algorithm/structure/tree/node"
)

func initBinaryTree() *Tree {
	tree := NewTree()
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

func createTree(num int) *Tree {
	tree := NewTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		r := rand.Intn(num)
		tree.Insert(r)
	}
	return tree
}

func TestInsert(t *testing.T) {
	num := 1000000
	tree := createTree(num)
	t.Log("插入节点数:", num)
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("旋转次数:", node.Count)
}

func TestDeleteLeaf(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(280)
	if tree.root.Left.Data == 160 &&
		tree.root.Right.Data == 240 &&
		tree.root.Right.Left.Data == 220 &&
		tree.root.Right.Right == nil {
		t.Log("删除280节点成功")
		return
	}
	t.Fatal("删除280节点失败")
}

func TestRightNode(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(160)
	if tree.root.Left.Data == 170 &&
		tree.root.Left.Right.Data == 180 &&
		tree.root.Left.Left.Data == 120 {
		t.Log("删除160节点成功")
		return
	}
	t.Fatal("删除160节点失败")
}

func TestRightNode1(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(200)
	if tree.root.Data == 220 &&
		tree.root.Left.Data == 160 &&
		tree.root.Right.Data == 240 &&
		tree.root.Right.Left == nil &&
		tree.root.Right.Right.Data == 280 {
		t.Log("删除200节点成功")
		return
	}
	t.Fatal("删除200节点失败")
}

func TestNotRightNode(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(80)
	if tree.root.Left.Left.Left.Data == 40 &&
		tree.root.Left.Left.Right.Data == 130 {
		t.Log("删除80节点成功")
		return
	}
	t.Fatal("删除80节点失败")
}

func TestNotRightNode1(t *testing.T) {
	tree := initBinaryTree()
	tree.Delete(180)
	if tree.root.Left.Right.Data == 170 &&
		tree.root.Left.Left.Data == 120 {
		t.Log("删除180节点成功")
		return
	}
	t.Fatal("删除180节点失败")
}

func TestNotRightNode2(t *testing.T) {
	tree := NewTree()
	tree.Insert(200)
	tree.Insert(160)
	tree.Insert(240)
	tree.Insert(120)
	tree.Insert(110)
	tree.Insert(140)
	tree.Insert(130)
	tree.Insert(150)
	tree.Insert(145)
	tree.Delete(160)
	if tree.root.Data == 200 &&
		tree.root.Left.Data == 150 &&
		tree.root.Right.Data == 240 &&
		tree.root.Left.Left.Data == 120 &&
		tree.root.Left.Left.Left.Data == 110 &&
		tree.root.Left.Left.Right.Data == 140 &&
		tree.root.Left.Left.Right.Left.Data == 130 &&
		tree.root.Left.Left.Right.Right.Data == 145 {
		t.Log("删除节点通过")
		return
	}
	t.Fatal("删除节点不通过")
}
