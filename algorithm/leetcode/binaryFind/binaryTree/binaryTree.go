package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		Val: data,
	}
}

var ret = make([]int, 0)

func inorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	return print(root)
}

func print(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left != nil {
		print(root.Left)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		print(root.Right)
	}
	return ret
}

func insert(root *TreeNode, data int) {
	if data < root.Val {
		if root.Left == nil {
			root.Left = NewTreeNode(data)
			return
		}
		insert(root.Left, data)
	}
	if data > root.Val {
		if root.Right == nil {
			root.Right = NewTreeNode(data)
			return
		}
		insert(root.Right, data)
	}
}
