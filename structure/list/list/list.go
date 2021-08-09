package list

type List struct {
	root    *Node
	pointer *Node
}

func NewList() *List {
	return &List{
		root: &Node{},
	}
}

func (l List) Append(data interface{}) {
	if l.root.next == nil {
		l.root.next = NewNode(data)
		return
	}
	l.root.next.Append(data)
}

func (l List) Remove(data interface{}) {
	if l.root.next == nil {
		return
	}
	if l.root.data == data {
		l.root.next = l.root.next.next
		return
	}
	l.root.Remove(data)
}

func (l List) RemoveAt(i int) {
	if l.root.next == nil {
		return
	}
	l.root.next.RemoveAt(i)
}

func (l List) Show(i int) interface{} {
	if l.root.next == nil {
		return nil
	}
	return l.root.next.Show(i)
}
