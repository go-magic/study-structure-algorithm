package tree

import "testing"

func TestLevelOrder3(t *testing.T) {
	tree := initTree3()
	trees := levelOrder3(tree)
	t.Log("测试通过", trees)
}

func initTree3() *TreeNode {
	tree := NewTreeNode(1)
	tree.Left = NewTreeNode(2)
	tree.Right = NewTreeNode(3)
	tree.Left.Left = NewTreeNode(4)
	tree.Right.Right = NewTreeNode(5)
	return tree
}

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
*/

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return printTree3(root)
}

func printTree3(root *TreeNode) [][]int {
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
		if len(rets)%2 == 1 {
			ret = reverse(ret)
		}
		rets = append(rets, ret)
	}
	return rets
}

func reverse(arr []int) []int {
	length := len(arr)
	ret := make([]int, length)
	for _, v := range arr {
		ret[length-1] = v
		length--
	}
	return ret
}
