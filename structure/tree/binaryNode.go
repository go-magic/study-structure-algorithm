package tree

import "fmt"

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
	return b.left
}

func (b *BinaryNode) Right() INode {
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
		return
	}
	if data > b.data {
		if b.right == nil {
			b.right = newBinaryNode(data)
			return
		}
		b.right.Insert(data)
		return
	}
}

/*
按层遍历二叉树的节点，要求不同层的节点需要换行
*/
func (b *BinaryNode) printRows() {
	//存储节点的临时队列
	tempBinaryNodes := make([]*BinaryNode, 0)
	//所有节点
	binaryNodes := make([]*BinaryNode, 0)
	//将节点加入队列
	tempBinaryNodes = append(tempBinaryNodes, b)
	binaryNodes = append(binaryNodes, b)
	binaryNodes = b.createQueue(tempBinaryNodes, binaryNodes)
	b.printRow(binaryNodes)
}

/*
把所有节点加入到队列中
*/
func (b *BinaryNode) createQueue(tempBinaryNodes, binaryNodes []*BinaryNode) []*BinaryNode {
	if len(tempBinaryNodes) == 0 {
		return binaryNodes
	}
	n := tempBinaryNodes[0]
	tempBinaryNodes = tempBinaryNodes[1:]
	if n.left != nil {
		tempBinaryNodes = append(tempBinaryNodes, n.left)
		binaryNodes = append(binaryNodes, n.left)
	}
	if n.right != nil {
		tempBinaryNodes = append(tempBinaryNodes, n.right)
		binaryNodes = append(binaryNodes, n.right)
	}
	return b.createQueue(tempBinaryNodes, binaryNodes)
}

/*
打印队列
*/
func (b *BinaryNode) printRow(binaryNodes []*BinaryNode) {
	if len(binaryNodes) == 0 {
		return
	}
	tempData := binaryNodes[0].data
	for _, node := range binaryNodes {
		if tempData > node.data {
			//换行
			fmt.Println()
		}
		b.printer.print(node)
		tempData = node.data
	}
}

/***************************************************************************************/
/*遍历end*/
/***************************************************************************************/

/*
删除节点
*/
func (b *BinaryNode) delete(data int) {
	if data < b.data {
		if b.left != nil {
			if data == b.left.data {
				b.left = b.left.deleteNode()
				return
			}
			b.left.delete(data)
		}
	}
	if data > b.data {
		if b.right != nil {
			if data == b.right.data {
				b.right = b.right.deleteNode()
				return
			}
			b.right.delete(data)
		}
	}
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
