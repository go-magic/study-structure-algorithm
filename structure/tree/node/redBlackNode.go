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
			//把父节点变为黑色统一处理
			if r.Color == Red {
				r.setBlackColor()
			}
			//父节点是黑色的则当前节点是平衡的
			return r
		}
		r.Left = r.Left.Insert(data)
	}
	if data > r.Data {
		if r.Right == nil {
			r.Right = NewRedBlackNode(data)
			//把父节点变为黑色统一处理
			if r.Color == Red {
				r.setBlackColor()
			}
			//父节点是黑色的则当前节点是平衡的
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
	r.Color = Black
	r.length++
}

/*
设置节点颜色并改变黑色节点的总长度
*/
func (r *RedBlackNode) setRedColor() {
	r.Color = Red
	r.length--
}

/*
检测红黑树是否平衡
*/
func (r *RedBlackNode) check() *RedBlackNode {
	return r.change()
}

/*
调整红黑树
*/
func (r *RedBlackNode) change() *RedBlackNode {
	if r.Left == nil || r.Right == nil {
		if r.Color == Red {
			r.setBlackColor()
		}
		//黑色节点直接就平衡，因为新插入的子节点肯定是红色
		return r
	}
	if !r.checkBalance() {
		return r.changeNode()
	}
	return r
}

/*
检测是否平衡
*/
func (r *RedBlackNode) checkBalance() bool {
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
