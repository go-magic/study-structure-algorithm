package tree

import "fmt"

type printer interface {
	print(b INode)
}

type rowPrint struct {
}

func (r rowPrint) print(n INode) {
	fmt.Printf("%d ", n.GetData())
}

type prettyPrint struct {
}

func (p prettyPrint) print(n INode) {
	length := Length(n)
	width := length * 3 * 2
	p.printSpace(0, width, n.GetData())
}

/*
控制节点位置
*/
func (p prettyPrint) printSpace(pos, width, data int) {
	for i := 0; i < width; i++ {
		if pos != i {
			fmt.Printf(" ")
		} else {
			fmt.Printf("%d", data)
		}
	}
}
