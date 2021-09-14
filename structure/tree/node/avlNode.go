package node

var balance bool
var Count int

type AvlNode struct {
	Left  *AvlNode
	Right *AvlNode
	Data  int
}

func NewAvlNode(data int) *AvlNode {
	balance = true
	return &AvlNode{
		Data: data,
	}
}

/*
平衡二叉树插入节点,插入后会改变节点的位置
*/
func (a *AvlNode) Insert(data int) *AvlNode {
	if data < a.Data {
		if a.Left == nil {
			a.Left = NewAvlNode(data)
			return a
		}
		a.Left = a.Left.Insert(data)
	}
	if data > a.Data {
		if a.Right == nil {
			a.Right = NewAvlNode(data)
			return a
		}
		a.Right = a.Right.Insert(data)
	}
	return a.check()
}

/*
调整节点
*/
func (a *AvlNode) check() *AvlNode {
	if !a.balance() && balance {
		balance = false
		Count++
		return a.change()
	}
	return a
}

/*
调整节点实体
*/
func (a *AvlNode) change() *AvlNode {
	if a.Left == nil {
		if a.Right.Right == nil {
			return a.rl()
		}
		return a.rr()
	}
	if a.Right == nil {
		if a.Left.Left == nil {
			return a.lr()
		}
		return a.ll()
	}
	if a.Left.Length() > a.Right.Length() {
		if a.Left.Left == nil {
			return a.lr()
		}
		return a.ll()
	}
	if a.Right.Right == nil {
		return a.rl()
	}
	return a.rr()
}

/*
左左
*/
func (a *AvlNode) ll() *AvlNode {
	l := a.Left
	a.Left = l.Right
	l.Right = a
	return l
}

/*
左右,把左左交给父节点旋转
*/
func (a *AvlNode) lr() *AvlNode {
	l := a.Left
	a.Left = l.Right
	l.Right = a.Left.Left
	a.Left.Left = l
	return a.ll()
}

/*
右左,把右右节点交给父节点旋转
*/
func (a *AvlNode) rl() *AvlNode {
	r := a.Right
	a.Right = r.Left
	r.Left = a.Right.Right
	a.Right.Right = r
	return a.rr()
}

/*
右右
*/
func (a *AvlNode) rr() *AvlNode {
	r := a.Right
	a.Right = r.Left
	r.Left = a
	return r
}

/*
检测节点是否平衡
*/
func (a *AvlNode) balance() bool {
	l, r := 0, 0
	if a.Left != nil {
		l = a.Left.Length() + 1
	}
	if a.Right != nil {
		r = a.Right.Length() + 1
	}
	if l > r {
		if l-r > 1 {
			return false
		}
	}
	if l < r {
		if r-l > 1 {
			return false
		}
	}
	return true
}

/*
获取树的层高
*/
func (a *AvlNode) Length() int {
	return a.length()
}

func (a *AvlNode) length() int {
	var l, r int
	if a.Left != nil {
		l = a.Left.length() + 1
	}
	if a.Right != nil {
		r = a.Right.length() + 1
	}
	if l > r {
		return l
	}
	return r
}

func (a *AvlNode) NodeNum() int {
	leftNum, rightNum := 0, 0
	if a.Left != nil {
		leftNum = a.Left.NodeNum() + 1
	}
	if a.Right != nil {
		rightNum = a.Right.NodeNum() + 1
	}
	return leftNum + rightNum
}
