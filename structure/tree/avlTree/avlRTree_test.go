package avl

import (
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
	tree.Insert(100)
	tree.Insert(40)
	return tree
}

func TestNewTree(t *testing.T) {
	insert()
}

func TestInsert(t *testing.T) {
	tree := createTree(100000)
	t.Log(tree.Length())
}

func TestNodeNum(t *testing.T) {
	tree := insert()
	t.Log(tree.NodeNum())
}
