package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Constructor()
	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	stack.Push(1)
	stack.Pop()
	stack.Pop()
	stack.Push(0)
}
