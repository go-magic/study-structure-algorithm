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
	for i := 0; i < length; i++ {
		bs := m[string(digits[i])]
		arrs = append(arrs, bs)
	}
	rets := Deep(arrs)
	return rets
}

func Deep(arrs [][]string) []string {
	return deep(0, arrs, []string{}, "")
}

func deep(x int, arrs [][]string, ret []string, s string) []string {
	if x == len(arrs) {
		ret = append(ret, s)
		return ret
	}
	if len(arrs[x]) == 0 {
		return ret
	}
	deep(x+1, arrs, ret, s+arrs[x][0])
	length := len(arrs[x])
	for i := 0; i < length; i++ {
		ret = deep(x+1, arrs, ret, s+arrs[x][i])
	}
	return ret
}

/*
[2,3,4] 深度优先

	a			b		c
   def		   def     def
   ghi         ghi     ghi
*/

func TestLetterCombinations(t *testing.T) {
	arr := letterCombinations1("78")
	t.Log(arr)
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}
	digit := string(digits[index])
	letters := phoneMap[digit]
	lettersCount := len(letters)
	for i := 0; i < lettersCount; i++ {
		backtrack(digits, index+1, combination+string(letters[i]))
	}
}
