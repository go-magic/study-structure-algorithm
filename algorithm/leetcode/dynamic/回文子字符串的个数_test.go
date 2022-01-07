package dynamic

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	if value := countSubstrings("abc"); value != 3 {
		t.Fatal("测试不通过:", value)
		return
	}
	if value := countSubstrings("aaa"); value != 6 {
		t.Fatal("测试不通过:", value)
		return
	}
	t.Log("测试通过")
}

func countSubstrings(s string) int {
	length := len(s)
	count := 0
	bytes := []byte(s)
	for i := length - 1; i >= 0; i-- {
		count += findMaxSubString(i, bytes)
	}
	return count
}

func findMaxSubString(index int, b []byte) int {
	length := len(b)
	count := 0
	for i := index; i < length; i++ {
		if isSubString(index, i, b) {
			count++
		}
	}
	return count
}

func isSubString(start, end int, b []byte) bool {
	for start <= end {
		if b[start] != b[end] {
			return false
		}
		start++
		end--
	}
	return true
}
