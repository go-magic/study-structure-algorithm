package dynamic

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	if num := numDecodings("12"); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("count:", count)
	count = 0
	if num := numDecodings("226"); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("count:", count)
	count = 0
	if num := numDecodings("0"); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("count:", count)
	count = 0
	if num := numDecodings("06"); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("count:", count)
	count = 0
	t.Log("测试通过")
}

func numDecodings(s string) int {
	numDecodingsLength = len(s)
	iniNumDecodingsSaver(numDecodingsLength)
	return getNumDecodings(0, 1, s) + getNumDecodings(0, 2, s)
}

var numDecodingsSaver []int
var numDecodingsLength int

func getNumDecodings(index int, step int, s string) int {
	count++
	now := index + step
	if now > numDecodingsLength {
		return 0
	}
	if now == numDecodingsLength {
		if ok := checNumDecodingskValid(index-1, step, s); !ok {
			return 0
		}
		return 1
	}
	if ok := checNumDecodingskValid(index-1, step, s); !ok {
		return 0
	}
	m := numDecodingsSaver[now]
	if m == -1 {
		m = getNumDecodings(index+step, 1, s) + getNumDecodings(index+step, 2, s)
		numDecodingsSaver[index+step] = m
	}
	return m
}

func iniNumDecodingsSaver(n int) {
	numDecodingsSaver = make([]int, 0, n)
	for i := 0; i < n; i++ {
		numDecodingsSaver = append(numDecodingsSaver, -1)
	}
}

func checNumDecodingskValid(index int, step int, s string) bool {
	now := index + step
	if s[index+1] == '0' {
		return false
	}
	if step == 1 {
		return true
	}
	if s[index+1] > '2' {
		return false
	}
	if s[index+1] == '2' && s[now] > '6' {
		return false
	}
	return true
}
