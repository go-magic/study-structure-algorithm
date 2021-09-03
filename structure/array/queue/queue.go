package queue

import "errors"

/*
循环数组实现队列
*/
type queue struct {
	pre    int
	next   int
	cap    int
	length int
	data   []interface{}
}

func NewQueue(cap int) *queue {
	return &queue{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

/*
Push 添加数据
*/
func (q *queue) Push(data interface{}) error {
	if q.length == q.cap {
		return errors.New("队列已满")
	}
	q.length++
	mod := q.mod(q.pre)
	q.data[mod] = data
	q.pre = q.advance(q.pre)
	return nil
}

/*
循环取余
*/
func (q *queue) advance(pre int) int {
	return (pre + 1) % q.cap
}

func (q *queue) mod(value int) int {
	return value % q.cap
}

/*
Pop 弹出数据
*/
func (q *queue) Pop() (interface{}, error) {
	if q.length == 0 {
		return nil, errors.New("当前队列没有任务")
	}
	q.length--
	mod := q.mod(q.next)
	data := q.data[mod]
	q.next = q.advance(q.next)
	return data, nil
}
