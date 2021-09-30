package list

import "testing"

func TestEnd(t *testing.T) {
	list := NewListNode(1)
	list.Next = NewListNode(3)
	list.Next.Next = NewListNode(2)
	arr := reversePrint1(list)
	t.Log(arr)
}

func TestReverse(t *testing.T) {
	list := NewListNode(1)
	list.appendList(2)
	list.appendList(3)
	list.appendList(4)
	list.appendList(5)
	p := reverseList(list)
	t.Log(p)
}

func TestReverse1(t *testing.T) {
	list := NewListNode(1)
	list.appendList(2)
	list.appendList(3)
	list.appendList(4)
	list.appendList(5)
	p := reverse1(list)
	t.Log(p)
}
