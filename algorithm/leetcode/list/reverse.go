package list

/*
反转链表
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverse(head)
}

/*
递归反转链表
*/
func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p := reverse(head.Next)
	p.appendList(head.Val)
	return p
}

/*
正向反转链表
*/
func reverse1(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p := head
	p1 := p.Next
	p2 := p.Next.Next
	for p2 != nil {
		p1.Next = p
		p = p1
		p1 = p2
		p2 = p2.Next
	}
	p1.Next = p
	head.Next = nil
	return p1
}
