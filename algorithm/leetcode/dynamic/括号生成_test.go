package dynamic

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	arr := generateParenthesis(4)
	t.Log(len(arr))
}

/*
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 2
输出：["()()","(())"]

输入：n = 1
输出：["()"]
*/

/*
分析: 输入n则有2*n个位子,只考虑插入‘(’，从2*n中选择n个插入。

*/
func generateParenthesis(n int) []string {
	return generateParenthesisDynamic(n)
}

/*
f(n) = f(插入'(') + f(插入')')


*/
func generateParenthesisDynamic(n int) []string {
	return generateDynamic(0, n, "(", "")
}

func generateDynamic(x, n int, s string, str string) []string {
	ret := make([]string, 0)
	if !checkValid(x, n, s, str) {
		return []string{}
	}
	str += s
	if x == n*2-1 {
		ret = append(ret, str)
		return ret
	}
	left := generateDynamic(x+1, n, "(", str)
	right := generateDynamic(x+1, n, ")", str)
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret
}

func checkValid(x, n int, s string, str string) bool {
	if x == n*2-1 && s == "(" {
		return false
	}
	num := findNum(s, str)
	if num == n {
		return false
	}
	if s == "(" {
		return true
	}
	left := findNum("(", str)
	right := findNum(")", str)
	return left > right
}

func findNum(s, str string) int {
	count := 0
	for _, v := range str {
		if s == string(v) {
			count++
		}
	}
	return count
}
