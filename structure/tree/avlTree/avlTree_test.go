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
	//tree.Insert(100)
	tree.Insert(40)
	return tree
}

func TestNewTree(t *testing.T) {
	insert()
}

func TestInsert(t *testing.T) {
	num := 10000000
	tree := createTree(num)
	t.Log("插入节点数:", num)
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("旋转次数:", node.Count)
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
