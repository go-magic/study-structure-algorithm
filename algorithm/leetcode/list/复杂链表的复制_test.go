package list

import "testing"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var mp = make(map[int]int, 0)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	target := copyNode(p)
	p1 := target
	i := 0
	for p != nil {
		setRandom(i, findRandom(head, p.Random))
		p = p.Next
		p1.Next = copyNode(p)
		p1 = p1.Next
		i++
	}
	p = head
	p1 = target
	i = 0
	for p != nil {
		p1.Random = findRandomAt(target, i)
		p = p.Next
		p1 = p1.Next
		i++
	}
	return target
}

func copyNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	return &Node{Val: n.Val}
}

func setRandom(val, i int) {
	mp[val] = i
}

func findRandomAt(head *Node, i int) *Node {
	at := mp[i]
	if at == -1 {
		return nil
	}
	for i := 0; i < at; i++ {
		head = head.Next
	}
	return head
}

func findRandom(head, p *Node) int {
	if p == nil {
		return -1
	}
	i := 0
	for head != nil {
		if head == p {
			return i
		}
		i++
		head = head.Next
	}
	return -1
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func insert(val int, n *Node) {
	if n == nil {
		return
	}
	for n.Next != nil {
		n = n.Next
	}
	n.Next = NewNode(val)
}

func insertRandom(i, j int, n *Node) {
	p := n
	for k := 0; k < i; k++ {
		p = p.Next
	}
	for k := 0; k < j; k++ {
		n = n.Next
	}
	p.Random = n
}

func initNode() *Node {
	n := NewNode(3)
	insert(5, n)
	insert(4, n)
	insert(-9, n)
	insert(-10, n)
	insert(5, n)
	insert(0, n)
	insert(6, n)
	insert(-6, n)
	insert(3, n)
	insert(-6, n)
	insert(9, n)
	insert(-2, n)
	insert(-3, n)
	insert(-1, n)
	insert(2, n)
	insert(-3, n)
	insert(-9, n)
	insert(-2, n)
	insert(-8, n)
	insert(5, n)

	//insertRandom(0, 17, n)
	insertRandom(1, 17, n)
	//insertRandom(2, 6, n)
	insertRandom(3, 6, n)
	insertRandom(4, 3, n)
	insertRandom(5, 15, n)
	insertRandom(6, 11, n)
	//insertRandom(7, 6, n)
	insertRandom(8, 16, n)
	insertRandom(9, 16, n)
	insertRandom(10, 11, n)
	insertRandom(11, 12, n)
	insertRandom(12, 1, n)
	insertRandom(13, 11, n)
	insertRandom(14, 10, n)
	insertRandom(15, 11, n)
	//insertRandom(16, 6, n)
	insertRandom(17, 7, n)
	insertRandom(18, 4, n)
	//insertRandom(19, 6, n)
	//insertRandom(20, 6, n)
	return n
}

func TestRandom(t *testing.T) {
	c := copyRandomList(initNode())
	t.Log(c)
}
