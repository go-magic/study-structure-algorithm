package redBlackTree

import (
	"github.com/go-magic/study-structure-algorithm/structure/tree/node"
	"math/rand"
	"testing"
	"time"
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
		tree.root.Left.Color == node.Black &&
		tree.root.Left.Data == 160 &&
		tree.root.Right.Color == node.Black &&
		tree.root.Right.Data == 240 &&
		tree.root.Left.Left.Color == node.Red &&
		tree.root.Left.Left.Data == 120 {
		t.Log("规则3通过")
		return
	}
	t.Fatal("规则3测试不通过")
}

func TestRule4(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(200)
	tree.Insert(160)
	tree.Insert(180)
	if tree.root.Color == node.Black &&
		tree.root.Data == 180 &&
		tree.root.Left.Color == node.Red &&
		tree.root.Left.Data == 160 &&
		tree.root.Right.Color == node.Red &&
		tree.root.Right.Data == 200 {
		t.Log("规则4通过")
		return
	}
	t.Fatal("规则4测试不通过")
}

func special() *Tree {
	tree := NewRedBlackTree()
	tree.Insert(13)
	tree.root.Left = node.NewRedBlackNode(8)
	tree.root.Left.Color = node.Black

	tree.root.Right = node.NewRedBlackNode(17)
	tree.root.Left.Left = node.NewRedBlackNode(1)
	tree.root.Left.Right = node.NewRedBlackNode(11)
	tree.root.Right.Left = node.NewRedBlackNode(15)
	tree.root.Right.Left.Color = node.Black
	tree.root.Right.Right = node.NewRedBlackNode(25)
	tree.root.Right.Right.Color = node.Black
	tree.root.Right.Right.Left = node.NewRedBlackNode(22)
	tree.root.Right.Right.Right = node.NewRedBlackNode(27)
	tree.Insert(21)
	return tree
}

//https://blog.csdn.net/bjweimengshu/article/details/106345677
func TestSpecial(t *testing.T) {
	tree := special()
	if tree.root.Data == 17 &&
		tree.root.Color == node.Black &&
		tree.root.Left.Data == 13 &&
		tree.root.Left.Color == node.Red &&
		tree.root.Right.Data == 25 &&
		tree.root.Right.Color == node.Red &&
		tree.root.Left.Left.Data == 8 &&
		tree.root.Left.Left.Color == node.Black &&
		tree.root.Left.Right.Data == 15 &&
		tree.root.Left.Right.Color == node.Black &&
		tree.root.Right.Left.Data == 22 &&
		tree.root.Right.Left.Color == node.Black &&
		tree.root.Right.Right.Data == 27 &&
		tree.root.Right.Right.Color == node.Black &&
		tree.root.Left.Left.Left.Data == 1 &&
		tree.root.Left.Left.Left.Color == node.Red &&
		tree.root.Left.Left.Right.Data == 11 &&
		tree.root.Left.Left.Right.Color == node.Red &&
		tree.root.Right.Left.Left.Data == 21 &&
		tree.root.Right.Left.Left.Color == node.Red {
		t.Log("手动组装的红黑树通过")
		t.Log("层高:", tree.Length())
		t.Log("节点数:", tree.NodeNum())
		t.Log("黑色节点数:", tree.root.Right.BlackLength())
		return
	}
	t.Fatal("手动组装的红黑树不通过")
}

//https://blog.csdn.net/xiaosong_2016/article/details/110954561
func TestSpecial1(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(10)
	if tree.root.Data != 10 || tree.root.Color != node.Black {
		t.Fatal("测试红黑树插入10失败")
		return
	}
	tree.Insert(5)
	if tree.root.Data != 10 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red {
		t.Fatal("测试红黑树插入5失败")
		return
	}
	tree.Insert(9)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red ||
		tree.root.Right.Data != 10 ||
		tree.root.Right.Color != node.Red {
		t.Fatal("测试红黑树插入9失败")
		return
	}
	tree.Insert(3)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Black ||
		tree.root.Right.Data != 10 ||
		tree.root.Right.Color != node.Black ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Red {
		t.Fatal("测试红黑树插入3失败")
		return
	}
	tree.Insert(6)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Black ||
		tree.root.Right.Data != 10 ||
		tree.root.Right.Color != node.Black ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Red ||
		tree.root.Left.Right.Data != 6 ||
		tree.root.Left.Right.Color != node.Red {
		t.Fatal("测试红黑树插入6失败")
		return
	}
	tree.Insert(7)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red ||
		tree.root.Right.Data != 10 ||
		tree.root.Right.Color != node.Black ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Black ||
		tree.root.Left.Right.Data != 6 ||
		tree.root.Left.Right.Color != node.Black ||
		tree.root.Left.Right.Right.Data != 7 ||
		tree.root.Left.Right.Right.Color != node.Red {
		t.Fatal("测试红黑树插入7失败")
		return
	}
	tree.Insert(19)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red ||
		tree.root.Right.Data != 10 ||
		tree.root.Right.Color != node.Black ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Black ||
		tree.root.Left.Right.Data != 6 ||
		tree.root.Left.Right.Color != node.Black ||
		tree.root.Left.Right.Right.Data != 7 ||
		tree.root.Left.Right.Right.Color != node.Red ||
		tree.root.Right.Right.Data != 19 ||
		tree.root.Right.Right.Color != node.Red {
		t.Fatal("测试红黑树插入19失败")
		return
	}
	tree.Insert(32)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red ||
		tree.root.Right.Data != 19 ||
		tree.root.Right.Color != node.Black ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Black ||
		tree.root.Left.Right.Data != 6 ||
		tree.root.Left.Right.Color != node.Black ||
		tree.root.Right.Left.Data != 10 ||
		tree.root.Right.Left.Color != node.Red ||
		tree.root.Right.Right.Data != 32 ||
		tree.root.Right.Right.Color != node.Red ||
		tree.root.Left.Right.Right.Data != 7 ||
		tree.root.Left.Right.Right.Color != node.Red {
		t.Fatal("测试红黑树插入32失败")
		return
	}
	tree.Insert(24)
	if tree.root.Data != 9 ||
		tree.root.Color != node.Black ||
		tree.root.Left.Data != 5 ||
		tree.root.Left.Color != node.Red ||
		tree.root.Right.Data != 19 ||
		tree.root.Right.Color != node.Red ||
		tree.root.Left.Left.Data != 3 ||
		tree.root.Left.Left.Color != node.Black ||
		tree.root.Left.Right.Data != 6 ||
		tree.root.Left.Right.Color != node.Black ||
		tree.root.Right.Left.Data != 10 ||
		tree.root.Right.Left.Color != node.Black ||
		tree.root.Right.Right.Data != 32 ||
		tree.root.Right.Right.Color != node.Black ||
		tree.root.Left.Right.Right.Data != 7 ||
		tree.root.Left.Right.Right.Color != node.Red ||
		tree.root.Right.Right.Left.Data != 24 ||
		tree.root.Right.Right.Left.Color != node.Red {
		t.Fatal("测试红黑树插入32失败")
		return
	}
	t.Log("测试红黑树成功")
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("黑色节点数:", tree.BlackLength())
}

func createTree(num int) *Tree {
	tree := NewRedBlackTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		r := rand.Intn(num)
		tree.Insert(r)
	}
	return tree
}

func TestSpecial2(t *testing.T) {
	num := 10000000
	tree := createTree(num)
	t.Log("插入节点数:", num)
	t.Log("层高:", tree.Length())
	t.Log("节点数:", tree.NodeNum())
	t.Log("旋转次数:", node.Count)
}
