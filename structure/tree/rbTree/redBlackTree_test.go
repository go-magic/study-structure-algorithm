package rbTree

import (
	"math/rand"
	"testing"
	"time"
)

func TestRule1(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	if tree.root.color == Red || tree.root.data != 10 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule2(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	tree.Insert(5)
	if tree.root.color == Red ||
		tree.root.data != 10 ||
		tree.root.Left.color == Black ||
		tree.root.Left.data != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule3(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(12)
	tree.Insert(1)
	if tree.root.color != Black && tree.root.data != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data != 5 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data != 1 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule4(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(2)
	if tree.root.color != Black && tree.root.data != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data != 2 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Red && tree.root.Left.Right.data != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule5(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(0)
	if tree.root.color != Black && tree.root.data != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data != 0 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Red && tree.root.Left.Right.data != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("层高:", tree.root.Length())
	t.Log("节点数量:", tree.root.Nodes()+1)
	t.Log("一条链上黑色节点数:", tree.root.BlackLength())
	t.Log("测试通过")
}

/*
程序员小灰案例测试
*/
func TestExample(t *testing.T) {
	tree := NewRedBlackTree()
	tree.root = NewRedBlackNode(13)
	tree.root.color = Black
	tree.root.Left = NewRedBlackNode(8)
	tree.root.Left.color = Black
	tree.root.Right = NewRedBlackNode(17)

	tree.root.Left.Left = NewRedBlackNode(1)
	tree.root.Left.Right = NewRedBlackNode(11)

	tree.root.Right.Left = NewRedBlackNode(15)
	tree.root.Right.Left.color = Black
	tree.root.Right.Right = NewRedBlackNode(25)
	tree.root.Right.Right.color = Black
	tree.root.Right.Right.Left = NewRedBlackNode(22)
	tree.root.Right.Right.Right = NewRedBlackNode(27)

	tree.Insert(21)
	if tree.root.color != Black && tree.root.data != 17 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Red && tree.root.Left.data != 13 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Red && tree.root.Right.data != 25 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Black && tree.root.Left.Left.data != 8 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Black && tree.root.Left.Right.data != 15 {
		t.Fatal("测试失败")
		return
	}

	if tree.root.Right.Left.color != Black && tree.root.Right.Left.data != 22 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.Right.color != Black && tree.root.Right.Right.data != 27 {
		t.Fatal("测试失败")
		return
	}

	if tree.root.Left.Left.Left.color != Red && tree.root.Left.Left.Left.data != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.Right.color != Red && tree.root.Left.Left.Right.data != 11 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.Left.Left.color != Red && tree.root.Right.Left.Left.data != 21 {
		t.Fatal("测试失败")
		return
	}
	t.Log("层高:", tree.root.Length())
	t.Log("节点数量:", tree.root.Nodes()+1)
	t.Log("一条链上黑色节点数:", tree.root.BlackLength())
	t.Log("测试通过")
}

func TestExample1(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(13)
	tree.root.Left = NewRedBlackNode(8)
	tree.root.Left.color = Black

	tree.root.Right = NewRedBlackNode(17)
	tree.root.Left.Left = NewRedBlackNode(1)
	tree.root.Left.Right = NewRedBlackNode(11)
	tree.root.Right.Left = NewRedBlackNode(15)
	tree.root.Right.Left.color = Black
	tree.root.Right.Right = NewRedBlackNode(25)
	tree.root.Right.Right.color = Black
	tree.root.Right.Right.Left = NewRedBlackNode(22)
	tree.root.Right.Right.Right = NewRedBlackNode(27)
	tree.Insert(21)

	t.Log("层高:", tree.root.Length())
	t.Log("节点数量:", tree.root.Nodes()+1)
	t.Log("一条链上黑色节点数:", tree.root.BlackLength())
	if tree.root.data == 17 &&
		tree.root.color == Black &&
		tree.root.Left.data == 13 &&
		tree.root.Left.color == Red &&
		tree.root.Right.data == 25 &&
		tree.root.Right.color == Red &&
		tree.root.Left.Left.data == 8 &&
		tree.root.Left.Left.color == Black &&
		tree.root.Left.Right.data == 15 &&
		tree.root.Left.Right.color == Black &&
		tree.root.Right.Left.data == 22 &&
		tree.root.Right.Left.color == Black &&
		tree.root.Right.Right.data == 27 &&
		tree.root.Right.Right.color == Black &&
		tree.root.Left.Left.Left.data == 1 &&
		tree.root.Left.Left.Left.color == Red &&
		tree.root.Left.Left.Right.data == 11 &&
		tree.root.Left.Left.Right.color == Red &&
		tree.root.Right.Left.Left.data == 21 &&
		tree.root.Right.Left.Left.color == Red {
		t.Log("手动组装的红黑树通过")
		return
	}
	t.Fatal("手动组装的红黑树不通过")
}

func createTree(num int) *RedBlackTree {
	tree := NewRedBlackTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		r := rand.Intn(num)
		tree.Insert(r)
	}
	return tree
}

func TestExample2(t *testing.T) {
	tree := createTree(10000000)
	//t.Log("层高:", tree.root.Length())
	//t.Log("节点数量:", tree.root.Nodes()+1)
	//t.Log("一条链上黑色节点数:", tree.root.BlackLength())
	t.Log(tree)
}
