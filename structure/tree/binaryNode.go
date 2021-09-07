package tree

type BinaryNode struct {
	left    *BinaryNode
	right   *BinaryNode
	data    int
	printer printer
}

/*
创建BinaryNode节点
*/
func newBinaryNode(data int) *BinaryNode {
	return &BinaryNode{data: data}
}

func (b *BinaryNode) GetData() int {
	return b.data
}

func (b *BinaryNode) Left() INode {
	//防止interface为空崩溃
	if b.left == nil {
		return nil
	}
	return b.left
}

func (b *BinaryNode) Right() INode {
	//防止interface为空崩溃
	if b.right == nil {
		return nil
	}
	return b.right
}

/*
插入节点
*/
func (b *BinaryNode) Insert(data int) {
	if data < b.data {
		if b.left == nil {
			b.left = newBinaryNode(data)
			return
		}
		b.left.Insert(data)
	}
	if data > b.data {
		if b.right == nil {
			b.right = newBinaryNode(data)
			return
		}
		b.right.Insert(data)
	}
}

/*
删除节点
*/
func (b *BinaryNode) Delete(data int) {
	if b.data == data {
		b.deleteRoot(data)
		return
	}
	b.delete(data)
}

func (b *BinaryNode) deleteRoot(data int) {
	left := b.left
	minNode := b.findNode(data)
	b.data = minNode.data
	minNode.left = left

}

func (b *BinaryNode) delete(data int) {

}

/*
删除节点
*/
func (b *BinaryNode) deleteNode() *BinaryNode {
	if b.left == nil && b.right == nil {
		return nil
	}
	if b.right == nil {

	}
	minNode := b.right.minNode()
	minNode.left = b.left
	p := minNode.right
	minNode.right = b.right
	b.right.left = p
	return minNode
}

/*
查找节点
*/
func (b *BinaryNode) findNode(data int) *BinaryNode {
	if data == b.data {
		return b
	}
	if data < b.data {
		return b.left.findNode(data)
	}
	if data > b.data {
		return b.right.findNode(data)
	}
	return nil
}

/*
查找最小阶段
*/
func (b *BinaryNode) minNode() *BinaryNode {
	if b.left == nil {
		return b
	}
	return b.left.minNode()
}
