package list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	p := head
	for head.Next != nil {
		addTail(p)
		head = head.Next
	}
	return p
}

func addTail(head *Node) {
	if head.Next == nil {
		head.Next = copyNode(head)
	}
	addTail(head.Next)
}

func copyNode(head *Node) *Node {
	return &Node{
		Val:    head.Val,
		Next:   head.Next,
		Random: head.Random,
	}
}
