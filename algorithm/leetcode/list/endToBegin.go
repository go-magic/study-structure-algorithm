package list

/*
倒序打印链表
*/

/*
递归解法
*/
func reversePrint(head *ListNode) []int {
	ret := make([]int, 0)
	if head == nil {
		return ret
	}
	if head.Next == nil {
		ret = append(ret, head.Val)
		return ret
	}
	ret = reversePrint(head.Next)
	ret = append(ret, head.Val)
	return ret
}

/*
循环解法
*/
func reversePrint1(head *ListNode) []int {
	temp := make([]int, 0)
	if head == nil {
		return temp
	}
	p := head
	for p.Next != nil {
		temp = append(temp, p.Val)
		p = p.Next
	}
	temp = append(temp, p.Val)
	ret := make([]int, len(temp), len(temp))
	pos := 0
	for i := len(temp) - 1; i >= 0; i-- {
		ret[i] = temp[pos]
		pos++
	}
	return ret
}
