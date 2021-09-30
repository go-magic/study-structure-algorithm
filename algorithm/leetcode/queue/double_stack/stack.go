package queue

type Stack struct {
	Data []int
}

func NewStack() *Stack {
	return &Stack{
		Data: make([]int, 0),
	}
}

func (s *Stack) Append(data int) {
	s.Data = append(s.Data, data)
}

func (s *Stack) Pop() int {
	length := len(s.Data)
	if length == 0 {
		return -1
	}
	data := s.Data[length-1]
	s.Data = s.Data[:length-1]
	return data
}

func (s *Stack) Length() int {
	return len(s.Data)
}
