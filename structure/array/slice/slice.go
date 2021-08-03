package slice

type Slice struct {
	data []interface{}
	len int
	cap int
}

/*
1M
*/
var mid = 1 << 20

/*
扩容策略
*/
func strategy(length int) int {
	if length < 1 {
		return 1
	}
	if length > mid {
		return int(float32(length) * 1.5)
	}
	return length << 1
}

/*
开辟一个新的空间
*/
func makeSlice(length int) []interface{} {
	return make([]interface{},0,length)
}

/*
Append 追加
*/
func (s *Slice)Append(data interface{})  {
	if s.len == s.cap {
		s.dilatation()
	}
	s.data = append(s.data,data)
	s.len++
}

/*
扩容
*/
func (s *Slice)dilatation()  {
	newCap := strategy(s.cap)
	slice := makeSlice(newCap)
	s.copy(&slice,s.data)
	s.data = slice
	s.cap = strategy(newCap)
}

/*
拷贝数据到新地址
*/
func (s *Slice)copy(target *[]interface{},source []interface{}){
	for i,_ := range source {
		*target = append(*target,source[i])
	}
}
