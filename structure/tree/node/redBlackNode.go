package node

type Color int

const (
	Red Color = iota
	Black
)

type RedBlackNode struct {
	Left   *RedBlackNode
	Right  *RedBlackNode
	Data   int
	Color  Color
	length int
}

/*
详情参考redBlackMap.go中的规则
*/

/*
默认插入红色节点
*/
func NewRedBlackNode(data int) *RedBlackNode {
	return &RedBlackNode{
		Data:  data,
		Color: Red,
	}
}

/*
插入节点
*/
func (r *RedBlackNode) Insert(data int) *RedBlackNode {
	if data < r.Data {
		if r.Left == nil {
			r.Left = NewRedBlackNode(data)
			return r
		}
		r.Left = r.Left.Insert(data)
	}
	if data > r.Data {
		if r.Right == nil {
			r.Right = NewRedBlackNode(data)
			return r
		}
		r.Right = r.Right.Insert(data)
	}
	return r.check()
}

/*
设置节点颜色并改变黑色节点的总长度
*/
func (r *RedBlackNode) setBlackColor() {
	if r.Color == Red {
		r.length++
	}
	r.Color = Black
}

/*
设置节点颜色并改变黑色节点的总长度
*/
func (r *RedBlackNode) setRedColor() {
	if r.Color == Black {
		r.length--
	}
	r.Color = Red
}

/*
检测红黑树是否平衡
*/
func (r *RedBlackNode) check() *RedBlackNode {
	if r.balance() {
		return r
	}
	return r.change()
}

/*
调整红黑树,规则1交给tree处理
*/
func (r *RedBlackNode) change() *RedBlackNode {
	//当前节点为红色则交由父节点处理
	if r.Color == Red {
		return r
	}
	if r.isRule2() {
		return r
	}
	if r.isRule3() {
		return r.rule3()
	}
	if r.isRule4() {
		return r.rule4()
	}
	if r.isRule5() {
		return r.rule5()
	}
	return r
}

/*
规则2
*/
func (r *RedBlackNode) isRule2() bool {
	//能走到这里，肯定有一个节点不为空
	if r.Left == nil && r.Right.Color == Red && r.Right.isLeaf() {
		return true
	}
	if r.Right == nil && r.Left.Color == Red && r.Left.isLeaf() {
		return true
	}
	return false
}

func (r *RedBlackNode) isRule3() bool {
	if r.Left != nil && r.Right != nil {
		//外部条件已经判断为这个节点为黑色，且不平衡，则左右孩子都为红色时则为规则3
		if r.Left.Color == Red && r.Right.Color == Red {
			if r.Left.haveRedChild() {
				return true
			}
			if r.Right.haveRedChild() {
				return true
			}
		}
	}
	return false
}

/*
规则4,节点的左孩子为红色，左孩子的右孩子为红色，右孩子不存在或右孩子为黑色。节点的右孩子为红色，右孩子的左孩子为红色，左孩子不存在或左孩子为黑色
*/
func (r *RedBlackNode) isRule4() bool {
	if r.Left == nil || r.Left.Color == Black {
		if r.Right.Left == nil {
			return false
		}
		if r.Right.Color == Red && r.Right.Left.Color == Red {
			return true
		}
		return false
	}
	if r.Right == nil || r.Right.Color == Black {
		if r.Left.Right == nil {
			return false
		}
		if r.Left.Color == Red && r.Left.Right.Color == Red {
			return true
		}
	}
	return false
}

/*
是否是规则5
*/
func (r *RedBlackNode) isRule5() bool {
	if r.Right == nil || r.Right.Color == Black {
		if r.Left.Color == Black || r.Left.Left == nil {
			return false
		}
		if r.Left.Color == Red && r.Left.Left.Color == Red {
			return true
		}
	}
	if r.Left == nil || r.Left.Color == Black {
		if r.Right.Color == Black || r.Right.Right == nil {
			return false
		}
		if r.Right.Color == Red && r.Right.Right.Color == Red {
			return true
		}
	}
	return false
}

func (r *RedBlackNode) rule3() *RedBlackNode {
	r.setRedColor()
	r.Left.setBlackColor()
	r.Right.setBlackColor()
	return r
}

/*
规则4
*/
func (r *RedBlackNode) rule4() *RedBlackNode {
	if r.Left != nil && r.Left.Color == Red {
		return r.lr()
	}
	return r.rl()
}

/*
规则5
*/
func (r *RedBlackNode) rule5() *RedBlackNode {
	if r.Left.Color == Red {
		return r.ll()
	}
	return r.rr()
}

/*
rl双旋转
*/
func (r *RedBlackNode) rl() *RedBlackNode {
	right := r.Right
	r.Right = right.Left
	right.Left = r.Right.Right
	r.Right.Right = right
	return r.rr()
}

/*
rr单旋转
*/
func (r *RedBlackNode) rr() *RedBlackNode {
	right := r.Right
	r.Right = right.Left
	right.Left = r
	right.setBlackColor()
	right.Left.setRedColor()
	return r
}

/*
lr双旋转
*/
func (r *RedBlackNode) lr() *RedBlackNode {
	left := r.Left
	r.Left = left.Right
	left.Right = r.Left.Left
	r.Left.Left = left
	return r.ll()
}

/*
ll单旋转
*/
func (r *RedBlackNode) ll() *RedBlackNode {
	left := r.Left
	r.Left = left.Right
	left.Right = r
	left.setBlackColor()
	left.Right.setRedColor()
	return left
}

/*
检测是否平衡
*/
func (r *RedBlackNode) balance() bool {
	if r.Color == Red && (r.Left.Color == Red || r.Right.Color == Red) {
		return false
	}
	return r.Left.length == r.Right.length
}

/*
调整分支
*/
func (r *RedBlackNode) changeNode() *RedBlackNode {
	if r.Color == Black {
		if r.Left.Color == Red {
			r.Left.setBlackColor()
		}
		if r.Right.Color == Red {
			r.Right.setBlackColor()
		}
		r.setRedColor()
		return r
	}
	return r.changeRedBranch()
}

/*
连续两个红色
*/
func (r *RedBlackNode) changeRedBranch() *RedBlackNode {
	if r.Left.Color == Red || r.Right.Color == Red {
		return r.changeBranch()
	}
	return r
}

func (r *RedBlackNode) changeBranch() *RedBlackNode {
	r.setBlackColor()
	return r
}

func (r *RedBlackNode) isLeaf() bool {
	if r.Left == nil && r.Right == nil {
		return true
	}
	return false
}

func (r *RedBlackNode) haveRedChild() bool {
	if r.Left != nil && r.Left.Color == Red {
		return true
	}
	if r.Right != nil && r.Right.Color == Red {
		return true
	}
	return false
}
