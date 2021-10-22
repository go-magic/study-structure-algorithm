package tree

import "testing"

func TestLevelOrders(t *testing.T) {
	tree := initTree1()
	trees := levelOrders(tree)
	t.Log("测试通过", trees)
}

func initTree1() *TreeNode {
	tree := NewTreeNode(-8)
	tree.Left = NewTreeNode(-6)
	tree.Right = NewTreeNode(7)
	tree.Left.Left = NewTreeNode(6)
	tree.Left.Left.Right = NewTreeNode(4)
	tree.Left.Left.Right.Left = NewTreeNode(5)
	return tree
}

func levelOrders(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return printTrees(root)
}

func printTrees(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	rets := make([][]int, 0)
	for len(queue) != 0 {
		length := len(queue)
		ret := make([]int, 0)
		for i := 0; i < length; i++ {
			n := queue[0]
			queue = queue[1:]
			ret = append(ret, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		rets = append(rets, ret)
	}
	return rets
}
