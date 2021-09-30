package list

/*
链表基本结构
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{
		Val: Val,
	}
}

func (l *ListNode) Append(Val int) {
	l.appendList(Val)
}

func (l *ListNode) appendList(Val int) {
	if l.Next == nil {
		l.Next = NewListNode(Val)
		return
	}
	l.Next.appendList(Val)
}

func (l *ListNode) appendNode(node *ListNode) {
	if l.Next == nil {
		l.Next = node
		return
	}
	l.Next.appendNode(node)
}

func (l *ListNode) appendListByLoop(Val int) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = NewListNode(Val)
}
