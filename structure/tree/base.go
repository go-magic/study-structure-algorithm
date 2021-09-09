package tree

func FindNode(data int, n INode) INode {
	if data == n.GetData() {
		return n
	}
	if data < n.GetData() {
		return FindNode(data, n.Left())
	}
	if data > n.GetData() {
		return FindNode(data, n.Right())
	}
	return nil
}

func FindMinNode(n INode) INode {
	if n.Left() == nil {
		return n
	}
	return FindMinNode(n.Left())
}
