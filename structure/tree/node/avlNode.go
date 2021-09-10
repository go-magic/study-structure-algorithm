package node

type AvlNode struct {
	Node
}

func (a *AvlNode) Insert(data int) {
	if data < a.Data {
		if a.Left == nil {
			a.Left = NewNode(data)
			return
		}
		a.Left.Insert(data)
	}
	if data > a.Data {
		if a.Right == nil {
			a.Right = NewNode(data)
			return
		}
		a.Right.Insert(data)
	}
	a.adjust()
}

func (a *AvlNode) adjust() {

}
