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

/*
				200                                   200                                              200
			   /   \
             160   240                           160        240                                    160     280
             / \   / \
           120 180220280     -----旋转----->   80    180  220     280   ------删除240---->        80   180 220
           /
          80                               40    120                                           40 120
         /
        40
*/
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
	//tree.Insert(130)
	//tree.Insert(100)
	tree.Insert(40)
	return tree
}

/*
删除叶子节点
*/
func TestDeleteLeafNode(t *testing.T) {
	tree := insert()
	tree.Delete(280)
	if tree.root.Right.Data == 240 &&
		tree.root.Right.Left.Data == 220 &&
		tree.root.Right.Right == nil {
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
	tree.Delete(120)
	if tree.root.Left.Left.Data == 80 &&
		tree.root.Left.Left.Left.Data == 40 &&
		tree.root.Left.Right == nil {
		t.Log("删除120成功")
		return
	}
	t.Fatal("删除120失败")
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
	tree.Delete(160)
	if tree.root.Left.Data == 80 &&
		tree.root.Left.Left.Data == 40 &&
		tree.root.Left.Right.Data == 180 &&
		tree.root.Left.Right.Left.Data == 120 {
		t.Log("删除160成功")
		return
	}
	t.Fatal("删除160失败")
}

func TestInsert(t *testing.T) {
	num := 1000000
	tree := createTree(num)
	t.Log("插入节点数:", num)
	t.Log("层高1:", tree.RecursionLength())
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("旋转次数:", node.Count)

	t.Log("开始删除节点")
	deleteNode(tree, 1000000, 10)
	t.Log("层高1:", tree.RecursionLength())
	t.Log("层高2:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
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
