package slice

import "testing"

func TestStrategy(t *testing.T)  {
	t.Log(strategy(2))
}

func TestSlice_Append(t *testing.T) {
	s := Slice{}
	s.Append(1)
	s.Append(2)
	s.Append(3)
	s.Append(4)
	s.Append(5)
	s.Append(6)
	s.Append(7)
	s.Append(8)
	t.Log(s)
}