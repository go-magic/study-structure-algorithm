package node

type Color int

const (
	Red Color = iota
	Black
)

type RedBlackNode struct {
	Left  *RedBlackNode
	Right *RedBlackNode
	Data  int
	Color Color
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
设置节点颜色
*/
func (r *RedBlackNode) SetBlackColor() {
	r.Color = Black
}

/*
设置节点颜色
*/
func (r *RedBlackNode) SetRedColor() {
	r.Color = Red
}

/*
检测红黑树是否平衡
*/
func (r *RedBlackNode) check() *RedBlackNode {
	if r.Left != nil {
		if !r.Left.balance() {
			return r.change()
		}
	}
	if r.Right != nil {
		if !r.Right.balance() {
			return r.change()
		}
	}
	return r
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
是否满足规则2
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

/*
是否满足规则3
*/
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
是否满足规则4,节点的左孩子为红色，左孩子的右孩子为红色，右孩子不存在或右孩子为黑色。节点的右孩子为红色，右孩子的左孩子为红色，左孩子不存在或左孩子为黑色
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
是否满足规则5
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

/*
规则3
*/
func (r *RedBlackNode) rule3() *RedBlackNode {
	r.SetRedColor()
	r.Left.SetBlackColor()
	r.Right.SetBlackColor()
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
	if r.Left != nil && r.Left.Color == Red {
		return r.ll()
	}
	return r.rr()
}

/*
rl双旋转,规则4
*/
func (r *RedBlackNode) rl() *RedBlackNode {
	right := r.Right
	r.Right = right.Left
	right.Left = r.Right.Right
	r.Right.Right = right
	return r.rr()
}

/*
rr单旋转,规则5
*/
func (r *RedBlackNode) rr() *RedBlackNode {
	right := r.Right
	r.Right = right.Left
	right.Left = r
	right.SetBlackColor()
	right.Left.SetRedColor()
	return right
}

/*
lr双旋转,规则4
*/
func (r *RedBlackNode) lr() *RedBlackNode {
	left := r.Left
	r.Left = left.Right
	left.Right = r.Left.Left
	r.Left.Left = left
	return r.ll()
}

/*
ll单旋转,规则5
*/
func (r *RedBlackNode) ll() *RedBlackNode {
	left := r.Left
	r.Left = left.Right
	left.Right = r
	left.SetBlackColor()
	left.Right.SetRedColor()
	return left
}

/*
检测是否平衡,每条枝干上的黑色节点数量为红黑树特性,判断是否平衡时不用这个特性,具体可以参考1-5规则
*/
func (r *RedBlackNode) balance() bool {
	if r.Color == Red {
		if r.Left != nil && r.Left.Color == Red {
			return false
		}
		if r.Right != nil && r.Right.Color == Red {
			return false
		}
	}
	return true
}

/*
是否是叶子节点
*/
func (r *RedBlackNode) isLeaf() bool {
	if r.Left == nil && r.Right == nil {
		return true
	}
	return false
}

/*
是否有红孩子节点
*/
func (r *RedBlackNode) haveRedChild() bool {
	if r.Left != nil && r.Left.Color == Red {
		return true
	}
	if r.Right != nil && r.Right.Color == Red {
		return true
	}
	return false
}

/*
获取层高
*/
func (r *RedBlackNode) Length() int {
	lLength, rLength := 0, 0
	if r.Left != nil {
		lLength = r.Left.Length() + 1
	}
	if r.Right != nil {
		rLength = r.Right.Length() + 1
	}
	if lLength > rLength {
		return lLength
	}
	return rLength
}

/*
获取黑色节点个数,根据红黑树特性判断1个子节点就可以了，这里判断了2个子节点
*/
func (r *RedBlackNode) BlackLength() int {
	lLength, rLength := 0, 0
	if r.Color == Black {
		lLength = 1
		rLength = 1
	}
	if r.Left != nil {
		lLength += r.Left.BlackLength()
	}
	if r.Right != nil {
		rLength += r.Right.BlackLength()
	}
	if lLength > rLength {
		return lLength
	}
	return rLength
}

/*
获取节点数量
*/
func (r *RedBlackNode) NodeNum() int {
	lNum, rNum := 0, 0
	if r.Left != nil {
		lNum += r.Left.NodeNum() + 1
	}
	if r.Right != nil {
		lNum += r.Right.NodeNum() + 1
	}
	return lNum + rNum
}

/*
删除节点
*/
func (r *RedBlackNode) DeleteNode(data int) *RedBlackNode {
	if data == r.Data {
		return r.deleteNode(data)
	}
	if data < r.Data {
		if r.Left != nil {
			r.Left = r.Left.DeleteNode(data)
		}
	}
	if data > r.Data {
		if r.Right != nil {
			r.Right = r.Right.DeleteNode(data)
		}
	}
	return r
}

func (r *RedBlackNode) deleteBalance() *RedBlackNode {

}

/*
删除节点实例
*/
func (r *RedBlackNode) deleteNode(data int) *RedBlackNode {
	if r.Right == nil {
		return r.Left
	}
	if r.Left == nil {
		return r.Right
	}
	return r.deleteExistRightNode()
}

/*
删除存在右节点的节点
*/
func (r *RedBlackNode) deleteExistRightNode() *RedBlackNode {
	minNode := r.Right.findMinNode()
	//minNode的left为空所以不用保存
	minNode.Left = r.Left
	//查找father是为了删除该节点
	father := r.Right.findFather(minNode.Data)
	if father == nil {
		return minNode
	}
	father.Left = minNode.Right
	minNode.Right = r.Right
	return minNode
}

/*
查找最小节点
*/
func (r *RedBlackNode) findMinNode() *RedBlackNode {
	if r.Left == nil {
		return r
	}
	return r.Left.findMinNode()
}

/*
查找最大节点
*/
func (r *RedBlackNode) findMaxNode() *RedBlackNode {
	if r.Right == nil {
		return r
	}
	return r.Right.findMaxNode()
}

/*
查找father
*/
func (r *RedBlackNode) findFather(data int) *RedBlackNode {
	if data < r.Data && r.Left != nil {
		if r.Left.Data == data {
			return r.Left
		}
		return r.Left.findFather(data)
	}
	if data > r.Data && r.Right != nil {
		if r.Right.Data == data {
			return r.Right
		}
		return r.Right.findFather(data)
	}
	return nil
}
