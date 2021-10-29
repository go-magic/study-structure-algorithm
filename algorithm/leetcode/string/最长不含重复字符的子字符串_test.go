package string

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	if str := lengthOfLongestSubstring("abcabcbb"); str != 3 {
		t.Fatal("测试失败")
		return
	}
	if str := lengthOfLongestSubstring("bbbbb"); str != 1 {
		t.Fatal("测试失败")
		return
	}
	if str := lengthOfLongestSubstring("pwwkew"); str != 3 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {

}
