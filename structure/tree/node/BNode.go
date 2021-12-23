package node

var M int

type Block struct {
	Head   *BNode
	Length int
}

type BNode struct {
	Data       int
	Pre        *BNode
	Next       *BNode
	LeftChild  *Block
	RightChild *Block
}

func NewBlock(data int) *Block {
	return &Block{
		Head:   NewBNode(data),
		Length: 1,
	}
}

func NewBNode(data int) *BNode {
	return &BNode{
		Data: data,
	}
}

func (b *Block) Insert(data int) {
	p := b.Head
	count := 0
	for p != nil {
		if data == p.Data {
			return
		}
		if data < p.Data {
			if p.LeftChild != nil {
				p.LeftChild.Insert(data)
				if p.LeftChild.Length == M {
					node := p.LeftChild.getMidNode()
					n := NewBNode(node.Data)
					n.Next = p
					n.Pre = p.Pre
					p.Pre = n
					n.LeftChild = p.LeftChild
					n.LeftChild.Length = M / 2
					p.LeftChild = &Block{Head: node}
					p.LeftChild.Length = M - M/2
					if count == 0 {
						b.Head = n
					}
					b.Length++
				}
				return
			}
			b.Length++
			node := NewBNode(data)
			if count == 0 {
				node.Next = b.Head
				b.Head = node
				node.Next.Pre = node
				return
			}
			p.Next = node
			node.Pre = p
			return
		}
		if p.Next == nil {
			if p.RightChild != nil {
				p.RightChild.Insert(data)
				if p.RightChild.Length == M {
					node := p.RightChild.getMidNode()
					p.Next = NewBNode(node.Data)
					p.Next.Pre = p
					p.Next.LeftChild = p.RightChild
					p.Next.LeftChild.Length = M / 2
					p.RightChild = nil
					p.Next.RightChild = &Block{Head: node}
					p.Next.RightChild.Length = M - M/2
					b.Length++
				}
				return
			}
			b.Length++
			p.Next = NewBNode(data)
			p.Next.Pre = p
			return
		}
		count++
		p = p.Next
	}
}

func (b *Block) getMidNode() *BNode {
	m := M / 2
	p := b.Head
	for i := 0; i < m; i++ {
		p = p.Next
	}
	return p
}
