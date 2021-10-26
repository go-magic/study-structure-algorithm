package rbTree

type Color int

const (
	Red Color = iota
	Black
)

type RedBlackNode struct {
	color       Color
	data        int
	Left, Right *RedBlackNode
}

/*
初始化为红色节点
*/
func NewRedBlackNode(data int) *RedBlackNode {
	return &RedBlackNode{
		color: Red,
		data:  data,
	}
}

/*
插入数据
*/
func (n *RedBlackNode) Insert(data int) *RedBlackNode {
	if data < n.data {
		if n.Left == nil {
			n.Left = NewRedBlackNode(data)
			return n
		}
		n.Left = n.Left.Insert(data)
	}
	if data > n.data {
		if n.Right == nil {
			n.Right = NewRedBlackNode(data)
			return n
		}
		n.Right = n.Right.Insert(data)
	}
	return n.check()
}

/*
检测红黑节点是否平衡
*/
func (n *RedBlackNode) check() *RedBlackNode {
	if n.Left != nil {
		if !n.Left.balance() {
			return n.change()
		}
	}
	if n.Right != nil {
		if !n.Right.balance() {
			return n.change()
		}
	}
	return n
}

/*
检测左节点是否平衡
*/
func (n *RedBlackNode) balance() bool {
	if n.color == Black {
		return true
	}
	if n.Left != nil && n.Left.color == Red {
		return false
	}
	if n.Right != nil && n.Right.color == Red {
		return false
	}
	return true
}

/*
调整红黑节点,规则1交给树处理，规则2是平衡的
*/
func (n *RedBlackNode) change() *RedBlackNode {
	if n.Left == nil || n.Left.color == Black {
		if n.Right.Right != nil && n.Right.Right.color == Red {
			return n.rule5R()
		}
		return n.rule4R()
	}
	if n.Right == nil || n.Right.color == Black {
		if n.Left.Left != nil && n.Left.Left.color == Red {
			return n.rule5L()
		}
		return n.rule4L()
	}
	return n.rule3()
}

/*
规则3,将左右孩子节点变为红色，自己变为黑色
*/
func (n *RedBlackNode) rule3() *RedBlackNode {
	n.Left.color = Black
	n.Right.color = Black
	n.color = Red
	return n
}

/*
规则4L,把规则4旋转成规则5
*/
func (n *RedBlackNode) rule4L() *RedBlackNode {
	node := n.Left
	n.Left = node.Right
	node.Right = n.Left.Left
	n.Left.Left = node
	return n.rule5L()
}

/*
规则4R,把规则4旋转成规则5
*/
func (n *RedBlackNode) rule4R() *RedBlackNode {
	node := n.Right
	n.Right = node.Left
	node.Left = n.Right.Right
	n.Right.Right = node
	return n.rule5R()
}

/*
规则5L,旋转规则5
*/
func (n *RedBlackNode) rule5L() *RedBlackNode {
	node := n.Left
	n.Left = node.Right
	node.Right = n
	node.Right.color = Red
	node.color = Black
	return node
}

/*
规则5R,旋转规则5
*/
func (n *RedBlackNode) rule5R() *RedBlackNode {
	node := n.Right
	n.Right = node.Left
	node.Left = n
	node.Left.color = Red
	node.color = Black
	return node
}

/*
获取节点数
*/
func (n *RedBlackNode) Nodes() int {
	left, right := 0, 0
	if n.Left != nil {
		left += n.Left.Nodes() + 1
	}
	if n.Right != nil {
		right += n.Right.Nodes() + 1
	}
	return left + right
}

/*
获取层高
*/
func (n *RedBlackNode) Length() int {
	left, right := 1, 1
	if n.Left != nil {
		left += n.Left.Length()
	}
	if n.Right != nil {
		right += n.Right.Length()
	}
	if left > right {
		return left
	}
	return right
}

/*
获取一条支线上的黑色节点数
*/
func (n *RedBlackNode) BlackLength() int {
	length := 0
	if n.color == Black {
		length = 1
	}
	if n.Left != nil {
		return length + n.Left.BlackLength()
	}
	return length
}
