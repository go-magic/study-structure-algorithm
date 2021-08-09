package list

/*
Node 递归实现
*/
type Node struct {
	data   interface{} /*数据*/
	next   *Node       /*后指针*/
	length int         /*链表长度*/
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) Append(data interface{}) {
	if n.next == nil {
		n.next = NewNode(data)
		n.length++
		return
	}
	n.next.Append(data)
}

func (n *Node) Remove(data interface{}) {
	if n.next != nil {
		if n.next.data == data {
			n.next = n.next.next
			n.length--
			return
		}
		n.next.Remove(data)
	}
}

func (n *Node) RemoveAt(i int) {
	if n.length < i+1 {
		return
	}
	n.removeAt(i, 0)
}

func (n *Node) removeAt(i, pos int) {
	if i == pos+1 {
		n.next = n.next.next
		n.length--
		return
	}
	n.removeAt(i, pos+1)
}

func (n *Node) Show(i int) interface{} {
	if n.length < i+1 {
		return nil
	}
	return n.show(i, 0)
}

func (n *Node) show(i, pos int) interface{} {
	if i == pos+1 {
		return n.next.data
	}
	return n.show(i, pos+1)
}
