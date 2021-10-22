package tree

import "testing"

func TestLevelOrder(t *testing.T) {
	tree := initTree()
	trees := levelOrder(tree)
	t.Log("测试通过", trees)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func initTree() *TreeNode {
	tree := NewTreeNode(3)
	tree.Left = NewTreeNode(9)
	tree.Right = NewTreeNode(20)
	tree.Right.Left = NewTreeNode(15)
	tree.Right.Right = NewTreeNode(7)
	return tree
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return printTree(root)
}

func printTree(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := make([]int, 0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		ret = append(ret, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return ret
}
