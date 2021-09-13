package node

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

/*
插入节点
*/
func (n *Node) Insert(data int) {
	if data < n.Data {
		if n.Left == nil {
			n.Left = NewNode(data)
			return
		}
		n.Left.Insert(data)
	}
	if data > n.Data {
		if n.Right == nil {
			n.Right = NewNode(data)
			return
		}
		n.Right.Insert(data)
	}
}

/*
寻找父节点
*/
func (n *Node) FindNodeFather(data int) *Node {
	if n.Left == nil && n.Right == nil {
		return nil
	}
	if n.Data == data {
		return nil
	}
	if n.Left != nil {
		if n.Left.Data == data {
			return n
		}
	}
	if n.Right != nil {
		if n.Right.Data == data {
			return n
		}
	}
	if data < n.Data {
		return n.Left.FindNodeFather(data)
	}
	return n.Right.FindNodeFather(data)
}

/*
查找最小节点
*/
func (n *Node) FindMinNode() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.FindMinNode()
}

/*
查找最大节点
*/
func (n *Node) FindMaxNode() *Node {
	if n.Right == nil {
		return n
	}
	return n.Right.FindMinNode()
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

func (n *Node) Delete(data int) *Node {
	return n.delete(data)
}

/*
删除节点实例
*/
func (n *Node) delete(data int) *Node {
	//删除当前节点
	if data == n.Data {
		if n.Left == nil && n.Right == nil {
			return nil
		}
		if n.Right == nil {
			return n.deleteRightNode()
		}
		return n.deleteLeftNode()
	}
	if data < n.Data {
		n.Left = n.Left.delete(data)
		return n
	}
	if data > n.Data {
		n.Right = n.Right.delete(data)
		return n
	}
	return nil
}

/*
删除右节点
*/
func (n *Node) deleteRightNode() *Node {
	maxNode := n.Left.FindMaxNode()
	father := n.Left.FindNodeFather(maxNode.Data)
	if father == nil {
		n.Left.Right = n.Left
		return n.Left
	}
	father.Right = maxNode.Left
	maxNode.Left = n.Left
	maxNode.Right = n.Right
	return maxNode
}

/*
删除左节点
*/
func (n *Node) deleteLeftNode() *Node {
	minNode := n.Right.FindMinNode()
	father := n.Right.FindNodeFather(minNode.Data)
	if father == nil {
		n.Right.Left = n.Left
		return n.Right
	}
	father.Left = minNode.Right
	minNode.Right = n.Right
	minNode.Left = n.Left
	return minNode
}

/*
获取树的层高
*/
func (n *Node) Length() int {
	return n.length()
}

func (n *Node) length() int {
	var l, r int
	if n.Left != nil {
		l = n.Left.length() + 1
	}
	if n.Right != nil {
		r = n.Right.length() + 1
	}
	if l > r {
		return l
	}
	return r
}
