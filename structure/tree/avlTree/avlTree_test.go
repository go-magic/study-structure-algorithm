package avl

import (
	"github.com/go-magic/study-structure-algorithm/structure/tree/node"
	"math/rand"
	"testing"
	"time"
)

func createTree(num int) *Tree {
	tree := NewTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		r := rand.Intn(num)
		tree.Insert(r)
	}
	return tree
}

func deleteNode(tree *Tree, length, num int) {
	for i := 0; i < num; i++ {
		r := rand.Intn(length)
		tree.Delete(r)
	}
}

func insert() *Tree {
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

/*
成功
*/
func TestInsertNode(t *testing.T) {
	tree := insert()
	if tree.root.Data == 160 &&
		tree.root.Left.Data == 120 &&
		tree.root.Right.Data == 200 &&
		tree.root.Right.Left.Data == 180 {
		t.Log("插入数据成功")
		return
	}
	t.Fatal("插入数据失败")
	return
}

/*
删除叶子节点
*/
func TestDeleteLeafNode(t *testing.T) {
	tree := insert()
	tree.Delete(280)
	if tree.root.Right.Right.Data == 240 &&
		tree.root.Right.Right.Left.Data == 220 &&
		tree.root.Right.Right.Right == nil {
		t.Log("删除280成功")
		return
	}
	t.Fatal("删除280失败")
}

/*
删除没有右孩子的节点
*/
func TestDeleteNotRightNode(t *testing.T) {
	tree := insert()
	tree.Delete(80)
	if tree.root.Left.Left.Data == 40 &&
		tree.root.Left.Data == 120 &&
		tree.root.Left.Right.Data == 130 &&
		tree.root.Left.Left.Right == nil {
		t.Log("删除80成功")
		return
	}
	t.Fatal("删除80失败")
}

/*
删除有右孩子的节点
*/
func deleteRightNode() bool {
	tree := insert()
	tree.Delete(240)
	if tree.root.Right.Data == 280 &&
		tree.root.Right.Left.Data == 220 &&
		tree.root.Right.Right == nil {
		return true
	}
	return false
}
func TestDeleteRightNode(t *testing.T) {
	tree := insert()
	tree.Delete(120)
	if tree.root.Left.Data == 80 &&
		tree.root.Left.Left.Data == 40 &&
		tree.root.Left.Right.Data == 130 &&
		tree.root.Right.Data == 200 {
		t.Log("删除120成功")
		return
	}
	t.Fatal("删除120失败")
}

func TestInsert(t *testing.T) {
	num := 10000000
	tree := createTree(num)
	t.Log("插入节点数:", num)
	t.Log("层高1:", tree.RecursionLength())
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("旋转次数:", node.Count)

	//t.Log("开始删除节点")
	//deleteNode(tree, 1000000, 10)
	//t.Log("层高1:", tree.RecursionLength())
	//t.Log("层高2:", tree.Length())
	//t.Log("节点数:", tree.NodeNum())
}

func TestNodeNum(t *testing.T) {
	tree := insert()
	t.Log(tree.NodeNum())
}

func makeSlice(data []int) {
	data = append(data, 1)
}

func TestSlice(t *testing.T) {
	data := make([]int, 0, 2)
	makeSlice(data)
	t.Log(data)
}
