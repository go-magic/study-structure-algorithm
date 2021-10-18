package string

import "testing"

var hash = map[string][]byte{}

func initMap() map[string][]string {
	m := make(map[string][]string)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	return m
}

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 {
		return []string{}
	}
	m := initMap()
	//base := m[string(digits[0])]
	arrs := make([][]string, 0, length)
	for i := 1; i < length; i++ {
		bs := m[string(digits[i])]
		arrs = append(arrs, bs)
	}
	rets := make([]string, 0)
	return rets
}

func TestLetterCombinations(t *testing.T) {
	arr := letterCombinations("23")
	t.Log(arr)
}

func initDeep() [][]string {
	return [][]string{{"d", "e", "f"}, {"g", "h", "i", "j"}}
}

func TestDeep(t *testing.T) {
	arrs := initDeep()
	arr := Deep(arrs)
	t.Log(arr)
}

func Deep(arrs [][]string) []string {
	return deep1(0, 0, arrs, []string{}, "")
}

func deep(x, y int, arrs [][]string, ret []string, s string) []string {
	if x == len(arrs) {
		ret = append(ret, s)
		return deep(0, y+1, arrs, ret, "")
	}
	if y == len(arrs[x]) {
		return ret
	}
	s += arrs[x][y]
	return deep(x+1, y, arrs, ret, s)
}

/*
[2,3,4] 深度优先

	a			b		c
   def		   def     def
   ghi         ghi     ghi
*/

func deep1(x, y int, arrs [][]string, ret []string, s string) []string {
	if x == len(arrs)-1 {
		ret = append(ret, s)
		return ret
	}
	if y == len(arrs[x]) {
		return ret
	}
	s += arrs[x][y]
	ret = deep1(x+1, y, arrs, ret, s)
	length := len(arrs)
	for i := 0; i < length; i++ {
		ret = deep1(x, i, arrs, ret, s)
	}
	return ret
}
