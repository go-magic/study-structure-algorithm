package list

import "testing"

func TestNewList(t *testing.T) {
	list := NewList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Remove(3)
	list.Remove(4)
	list.Remove(1)
	list.Append(5)
	list.Append(6)
}
