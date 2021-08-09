package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue(3)
	t.Log(q.Push(1))
	t.Log(q.Push(2))
	t.Log(q.Push(3))
	t.Log(q.Push(4))
	t.Log(q.Push(5))
	t.Log(q.Push(6))
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Push(7))
	t.Log(q.Push(8))
	t.Log(q.Push(9))
	t.Log(q.Push(10))
	t.Log(q.Push(11))
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Push(12))
	t.Log(q.Push(13))
	t.Log(q)
}
