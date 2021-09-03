package tree

type BinaryTree struct {
	root *BinaryNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

/*
插入数据
*/
func (b *BinaryTree) Insert(data int) {
	if b.root == nil {
		b.root = newBinaryNode(data)
		return
	}
	b.root.Insert(data)
}

/*
打印数据
*/
func (b *BinaryTree) Delete(data int) {
	if b.root == nil {
		return
	}
	if b.root.data == data {
		b.deleteRoot()
		return
	}
	b.root.delete(data)
}

/*
删除根节点
*/
func (b *BinaryTree) deleteRoot() {
	if b.root.left == nil {
		b.root = b.root.right
		return
	}
	if b.root.right == nil {
		b.root = b.root.left
		return
	}
}
