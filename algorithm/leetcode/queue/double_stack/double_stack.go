package queue

type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.stack1.Append(value)
}

func (q *CQueue) DeleteHead() int {
	if q.stack2.Length() != 0 {
		return q.stack2.Pop()
	}
	if q.stack1.Length() != 0 {
		q.move()
	}
	return q.stack2.Pop()
}

func (q *CQueue) move() {
	if q.stack1.Length() != 0 {
		data := q.stack1.Pop()
		for data != -1 {
			q.stack2.Append(data)
			data = q.stack1.Pop()
		}
	}
}
