package dynamic

import (
	"testing"
)

func TestSubGenerateParenthesis(t *testing.T) {
	strs := subGenerateParenthesis(3)
	t.Log(strs)
	strs = subGenerateParenthesis(1)
	t.Log(strs)
}

func subGenerateParenthesis(n int) []string {
	return subGenerateParenthesisString("", '(', 1, 1, n*2)
}

func subGenerateParenthesisString(s string, c byte, flag, index, n int) []string {
	if flag < 0 && c == ')' {
		return []string{}
	}
	if flag > n-index {
		return []string{}
	}
	if index == n {
		return []string{s + string(c)}
	}
	right := subGenerateParenthesisString(s+string(c), ')', flag-1, index+1, n)
	left := subGenerateParenthesisString(s+string(c), '(', flag+1, index+1, n)
	right = append(right, left...)
	return right
}

func dynamicSubGenerateParenthesis(n int) int {
	return dynamicSubGenerateParenthesisSub('(', 1, 1, n)
}

func dynamicSubGenerateParenthesisSub(c byte, flag, index, n int) int {
	if flag < 0 && c == ')' {
		return 0
	}
	if flag > n-index {
		return 0
	}
	if index == n {
		return 1
	}
	right := dynamicSubGenerateParenthesisSub(')', flag-1, index+1, n)
	left := dynamicSubGenerateParenthesisSub('(', flag+1, index+1, n)
	return right + left
}

func TestDynamicSubGenerateParenthesis(t *testing.T) {
	strs := dynamicSubGenerateParenthesis(2)
	t.Log(strs)
	strs = dynamicSubGenerateParenthesis(4)
	t.Log(strs)
	strs = dynamicSubGenerateParenthesis(6)
	t.Log(strs)
}
