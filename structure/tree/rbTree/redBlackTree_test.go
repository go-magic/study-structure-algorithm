package rbTree

import (
	"testing"
)

func TestRule1(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(10))
	if tree.root.color == Red || tree.root.data.Hash() != 10 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule2(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(10))
	tree.Insert(DefaultData(5))
	if tree.root.color == Red ||
		tree.root.data.Hash() != 10 ||
		tree.root.Left.color == Black ||
		tree.root.Left.data.Hash() != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule3(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(10))
	tree.Insert(DefaultData(5))
	tree.Insert(DefaultData(12))
	tree.Insert(DefaultData(1))
	if tree.root.color != Black && tree.root.data.Hash() != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data.Hash() != 5 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data.Hash() != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data.Hash() != 1 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule4(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(10))
	tree.Insert(DefaultData(5))
	tree.Insert(DefaultData(12))
	tree.Insert(DefaultData(1))
	tree.Insert(DefaultData(2))
	if tree.root.color != Black && tree.root.data.Hash() != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data.Hash() != 2 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data.Hash() != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data.Hash() != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Red && tree.root.Left.Right.data.Hash() != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestRule5(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(10))
	tree.Insert(DefaultData(5))
	tree.Insert(DefaultData(12))
	tree.Insert(DefaultData(1))
	tree.Insert(DefaultData(0))
	if tree.root.color != Black && tree.root.data.Hash() != 10 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Black && tree.root.Left.data.Hash() != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Black && tree.root.Right.data.Hash() != 12 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Red && tree.root.Left.Left.data.Hash() != 0 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Red && tree.root.Left.Right.data.Hash() != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

/*
程序员小灰案例测试
*/
func TestExample(t *testing.T) {
	tree := NewRedBlackTree()
	tree.root = NewRedBlackNode(DefaultData(13))
	tree.root.color = Black
	tree.root.Left = NewRedBlackNode(DefaultData(8))
	tree.root.Left.color = Black
	tree.root.Right = NewRedBlackNode(DefaultData(17))

	tree.root.Left.Left = NewRedBlackNode(DefaultData(1))
	tree.root.Left.Right = NewRedBlackNode(DefaultData(11))

	tree.root.Right.Left = NewRedBlackNode(DefaultData(15))
	tree.root.Right.Left.color = Black
	tree.root.Right.Right = NewRedBlackNode(DefaultData(25))
	tree.root.Right.Right.color = Black
	tree.root.Right.Right.Left = NewRedBlackNode(DefaultData(22))
	tree.root.Right.Right.Right = NewRedBlackNode(DefaultData(27))

	tree.Insert(DefaultData(21))
	if tree.root.color != Black && tree.root.data.Hash() != 17 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.color != Red && tree.root.Left.data.Hash() != 13 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.color != Red && tree.root.Right.data.Hash() != 25 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.color != Black && tree.root.Left.Left.data.Hash() != 8 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Right.color != Black && tree.root.Left.Right.data.Hash() != 15 {
		t.Fatal("测试失败")
		return
	}

	if tree.root.Right.Left.color != Black && tree.root.Right.Left.data.Hash() != 22 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.Right.color != Black && tree.root.Right.Right.data.Hash() != 27 {
		t.Fatal("测试失败")
		return
	}

	if tree.root.Left.Left.Left.color != Red && tree.root.Left.Left.Left.data.Hash() != 1 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Left.Left.Right.color != Red && tree.root.Left.Left.Right.data.Hash() != 11 {
		t.Fatal("测试失败")
		return
	}
	if tree.root.Right.Left.Left.color != Red && tree.root.Right.Left.Left.data.Hash() != 21 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestExample1(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(DefaultData(13))
	tree.root.Left = NewRedBlackNode(DefaultData(8))
	tree.root.Left.color = Black

	tree.root.Right = NewRedBlackNode(DefaultData(17))
	tree.root.Left.Left = NewRedBlackNode(DefaultData(1))
	tree.root.Left.Right = NewRedBlackNode(DefaultData(11))
	tree.root.Right.Left = NewRedBlackNode(DefaultData(15))
	tree.root.Right.Left.color = Black
	tree.root.Right.Right = NewRedBlackNode(DefaultData(25))
	tree.root.Right.Right.color = Black
	tree.root.Right.Right.Left = NewRedBlackNode(DefaultData(22))
	tree.root.Right.Right.Right = NewRedBlackNode(DefaultData(27))
	tree.Insert(DefaultData(21))

	if tree.root.data.Hash() == 17 &&
		tree.root.color == Black &&
		tree.root.Left.data.Hash() == 13 &&
		tree.root.Left.color == Red &&
		tree.root.Right.data.Hash() == 25 &&
		tree.root.Right.color == Red &&
		tree.root.Left.Left.data.Hash() == 8 &&
		tree.root.Left.Left.color == Black &&
		tree.root.Left.Right.data.Hash() == 15 &&
		tree.root.Left.Right.color == Black &&
		tree.root.Right.Left.data.Hash() == 22 &&
		tree.root.Right.Left.color == Black &&
		tree.root.Right.Right.data.Hash() == 27 &&
		tree.root.Right.Right.color == Black &&
		tree.root.Left.Left.Left.data.Hash() == 1 &&
		tree.root.Left.Left.Left.color == Red &&
		tree.root.Left.Left.Right.data.Hash() == 11 &&
		tree.root.Left.Left.Right.color == Red &&
		tree.root.Right.Left.Left.data.Hash() == 21 &&
		tree.root.Right.Left.Left.color == Red {
		t.Log("手动组装的红黑树通过")
		return
	}
	t.Fatal("手动组装的红黑树不通过")
}
