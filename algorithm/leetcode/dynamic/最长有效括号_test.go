package dynamic

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	if num := longestValidParentheses("(()"); num != 2 {
		t.Fatal("测试失败")
		return
	}
	t.Log(count1)
	count1 = 0
	if num := longestValidParentheses(")()())"); num != 4 {
		t.Fatal("测试失败")
		return
	}
	t.Log(count1)
	count1 = 0
	if num := longestValidParentheses(""); num != 0 {
		t.Fatal("测试失败")
		return
	}
	t.Log(count1)
	count1 = 0
	if num := longestValidParentheses("(()())"); num != 6 {
		t.Fatal("测试失败")
		return
	}
	t.Log(count1)
	count1 = 0
	t.Log("测试成功")
}

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0


提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

*/
func longestValidParentheses(s string) int {
	return longestValid(0, 0, "", s)
}

var count1 int

func longestValid(length, index int, s, str string) int {
	ok, valid := checkVlid(len(str), s)
	if !valid {
		return length
	}
	if ok {
		length = len(s)
	}
	if index == len(str) {
		return length
	}
	count1 += 2
	l := longestValid(length, index+1, s+string(str[index]), str)
	r := longestValid(0, index+1, "", str)
	if l > r {
		return l
	}
	return r
}

func checkVlid(n int, s string) (bool, bool) {
	left, right := 0, 0
	for _, v := range s {
		if v == '(' {
			left++
		} else {
			right++
		}
		if right > left || left > n/2 {
			return false, false
		}
	}
	return left == right, true
}
