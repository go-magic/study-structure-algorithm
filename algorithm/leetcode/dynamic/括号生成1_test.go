package dynamic

import (
	"testing"
)

func TestGenerateParenthesis1(t *testing.T) {
	arr := generateParenthesis1(1)
	t.Log(len(arr))
}

/*
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]
*/

/*
分析: 输入n则有2*n个位子,只考虑插入‘(’，从2*n中选择n个插入。

*/

func generateParenthesis1(n int) []string {
	return generateParenthesis2(1, 0, n, "(")
}

func generateParenthesis2(left, right, n int, str string) []string {
	ret := make([]string, 0)
	if left > n {
		return ret
	}
	if right == n {
		ret = append(ret, str)
		return ret
	}
	l := generateParenthesis2(left+1, right, n, str+"(")
	ret = append(ret, l...)
	if left > right {
		r := generateParenthesis2(left, right+1, n, str+")")
		ret = append(ret, r...)
	}
	return ret
}
