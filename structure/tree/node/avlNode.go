package node

var Count int

type AvlNode struct {
	Left       *AvlNode
	Right      *AvlNode
	Data       int
	NodeLength int
}

func NewAvlNode(data int) *AvlNode {
	return &AvlNode{
		Data:       data,
		NodeLength: 1,
	}
}

/*
平衡二叉树插入节点,插入后会改变节点的位置
*/
func (a *AvlNode) Insert(data int) *AvlNode {
	if data < a.Data {
		if a.Left == nil {
			a.Left = NewAvlNode(data)
			a.setNodeLength()
			return a
		}
		a.Left = a.Left.Insert(data)
	}
	if data > a.Data {
		if a.Right == nil {
			a.Right = NewAvlNode(data)
			a.setNodeLength()
			return a
		}
		a.Right = a.Right.Insert(data)
	}
	a.setNodeLength()
	return a.check()
}

/*
记录每个节点的层高,不用每次去递归获取,每次给左右节点赋值都需要重新计算层高
*/
func (a *AvlNode) setNodeLength() {
	l, r := 1, 1
	if a.Left != nil {
		l = a.Left.NodeLength + 1
	}
	if a.Right != nil {
		r = a.Right.NodeLength + 1
	}
	if l > r {
		a.NodeLength = l
		return
	}
	a.NodeLength = r
}

/*
调整节点
*/
func (a *AvlNode) check() *AvlNode {
	if !a.balance() {
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
	Count++
	l := a.Left
	a.Left = l.Right
	a.setNodeLength()
	l.Right = a
	l.setNodeLength()
	return l
}

/*
左右
*/
func (a *AvlNode) lr() *AvlNode {
	Count++
	l := a.Left
	a.Left = l.Right
	a.setNodeLength()
	l.Right = a.Left.Left
	l.setNodeLength()
	a.Left.Left = l
	a.Left.setNodeLength()
	return a.ll()
}

/*
右左
*/
func (a *AvlNode) rl() *AvlNode {
	Count++
	r := a.Right
	a.Right = r.Left
	a.setNodeLength()
	r.Left = a.Right.Right
	r.setNodeLength()
	a.Right.Right = r
	a.Right.setNodeLength()
	return a.rr()
}

/*
右右
*/
func (a *AvlNode) rr() *AvlNode {
	Count++
	r := a.Right
	a.Right = r.Left
	a.setNodeLength()
	r.Left = a
	r.setNodeLength()
	return r
}

/*
检测节点是否平衡
*/
func (a *AvlNode) balance() bool {
	l, r := 1, 1
	if a.Left != nil {
		l = a.Left.NodeLength + 1
	}
	if a.Right != nil {
		r = a.Right.NodeLength + 1
	}
	if l-r > 1 {
		return false
	}
	if r-l > 1 {
		return false
	}
	return true
}

func (a *AvlNode) RecursionLength() int {
	l, r := 0, 0
	if a.Left != nil {
		l = a.Left.RecursionLength() + 1
	}
	if a.Right != nil {
		r = a.Right.RecursionLength() + 1
	}
	if l > r {
		return l
	}
	return r
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
		l = a.Left.NodeLength + 1
	}
	if a.Right != nil {
		r = a.Right.NodeLength + 1
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

func (a *AvlNode) DeleteNode(data int) *AvlNode {
	return a.deleteNode(data)
}

func (a *AvlNode) deleteNode(data int) *AvlNode {
	if data < a.Data {
		if a.Left != nil {
			a.Left = a.Left.deleteNode(data)
			a.setNodeLength()
			return a.check()
		}
	}
	if data > a.Data {
		if a.Right != nil {
			a.Right = a.Right.deleteNode(data)
			a.setNodeLength()
			return a.check()
		}
	}
	return a.delete()
}

/*
二叉树删除涉及到多种情况，需要逐个处理

1.当前节点为叶子节点

　　直接删除

2.当前节点右子树为空

　　复制左子树中最大的值，用该值替代当前节点，删除左子树中原节点。

3.当前节点右子树不为空

　　复制右子树中最小的值，用该值替代当前节点，删除右子树中原节点。
*/

func (a *AvlNode) delete() *AvlNode {
	if a.Left == nil && a.Right == nil {
		return nil
	}
	if a.Right == nil {
		return a.deleteLeftNode()
	}
	return a.deleteRightNode()
}

/*
如果没有右节点那么左节点最多只有一个
*/
func (a *AvlNode) deleteLeftNode() *AvlNode {
	return a.Left
}

/*
把右子树的最小节点移上来,注意需要把最小节点的右节点挂到最小节点的父节点的左孩子上,注意要检查这个节点是否平衡
*/
func (a *AvlNode) deleteRightNode() *AvlNode {
	minNode := a.Right.findMinNode()
	minNode.Left = a.Left
	minNode.setNodeLength()
	father := a.Right.findNodeFather(minNode.Data)
	if father == nil {
		return minNode.check()
	}
	father.Left = minNode.Right
	father.setNodeLength()
	minNode.Right = a.Right
	minNode.setNodeLength()
	return minNode.check()
}

/*
查找最小节点
*/
func (a *AvlNode) findMinNode() *AvlNode {
	if a.Left == nil {
		return a
	}
	return a.Left.findMinNode()
}

/*
查找最大节点
*/
func (a *AvlNode) findMaxNode() *AvlNode {
	if a.Right == nil {
		return a
	}
	return a.Right.findMaxNode()
}

/*
寻找父节点
*/
func (a *AvlNode) findNodeFather(data int) *AvlNode {
	if a.Left == nil && a.Right == nil {
		return nil
	}
	if a.Data == data {
		return nil
	}
	if a.Left != nil {
		if a.Left.Data == data {
			return a
		}
	}
	if a.Right != nil {
		if a.Right.Data == data {
			return a
		}
	}
	if data < a.Data {
		return a.Left.findNodeFather(data)
	}
	return a.Right.findNodeFather(data)
}
