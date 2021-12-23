package b_Tree

import (
	"testing"
)

func TestName(t *testing.T) {
	M = 5
	tree := NewTree(5)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(12)
	tree.Insert(8)
}
