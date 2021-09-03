package tree

import "fmt"

type Tree struct {
	n       INode
	printer printer
}

func NewTree(n INode, printer printer) *Tree {
	return &Tree{
		n:       n,
		printer: printer,
	}
}

/*
查找最小节点
*/
func (t *Tree) FindMinNode() INode {
	if t.n == nil {
		return nil
	}
	return t.findMinNode(t.n)
}

func (t *Tree) findMinNode(n INode) INode {
	left := n.Left()
	if left == nil {
		return n
	}
	return t.findMinNode(left)
}

/*
前序遍历
*/
func (t *Tree) PrintPre() {
	t.printPre(t.n)
}

func (t *Tree) printPre(n INode) {
	fmt.Println(n.GetData())
	if n.Left() != nil {
		t.printPre(n.Left())
	}
	if n.Right() != nil {
		t.printPre(n.Right())
	}
}

/*
中序遍历
*/
func (t *Tree) PrintMid() {
	t.printMid(t.n)
}

func (t *Tree) printMid(n INode) {
	if n.Left() != nil {
		t.printMid(n.Left())
	}
	fmt.Println(n.GetData())
	if n.Right() != nil {
		t.printMid(n.Right())
	}
}

/*
中序遍历
*/
func (t *Tree) PrintNext() {
	t.printNext(t.n)
}

func (t *Tree) printNext(n INode) {
	if n.Left() != nil {
		t.printNext(n.Left())
	}
	if n.Right() != nil {
		t.printNext(n.Right())
	}
	fmt.Println(n.GetData())
}

/*
删除节点
*/
func (t *Tree) Delete(data int) {
	t.n.Delete(data)
}

/*
插入节点
*/
func (t *Tree) Insert(data int) {
	t.n.Insert(data)
}

/*
按层遍历二叉树的节点，要求不同层的节点需要换行
*/
func (t *Tree) PrintRows() {
	//存储节点的临时队列
	tempNodes := make([]INode, 0)
	//所有节点
	nodes := make([]INode, 0)
	//将节点加入队列
	tempNodes = append(tempNodes, t.n)
	nodes = append(nodes, t.n)
	nodes = t.createQueue(tempNodes, nodes)
	t.printRow(nodes)
}

/*
把所有节点加入到队列中
*/
func (t *Tree) createQueue(tempNodes, nodes []INode) []INode {
	if len(tempNodes) == 0 {
		return nodes
	}
	n := tempNodes[0]
	tempNodes = tempNodes[1:]
	if n.Left() != nil {
		tempNodes = append(tempNodes, n.Left())
		nodes = append(nodes, n.Left())
	}
	if n.Right() != nil {
		tempNodes = append(tempNodes, n.Right())
		nodes = append(nodes, n.Right())
	}
	return t.createQueue(tempNodes, nodes)
}

/*
打印队列
*/
func (t *Tree) printRow(binaryNodes []INode) {
	if len(binaryNodes) == 0 {
		return
	}
	tempData := binaryNodes[0].GetData()
	for _, node := range binaryNodes {
		if tempData > node.GetData() {
			//换行
			fmt.Println()
		}
		t.printer.print(node)
		tempData = node.GetData()
	}
}
