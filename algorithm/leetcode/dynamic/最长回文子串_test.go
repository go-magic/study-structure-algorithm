package dynamic

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	if str := longestPalindrome("babad"); str != "bab" {
		t.Fatal("测试失败")
		return
	}
	if str := longestPalindrome("cbbd"); str != "bb" {
		t.Fatal("测试失败")
		return
	}
	if str := longestPalindrome("a"); str != "a" {
		t.Fatal("测试失败")
		return
	}
	if str := longestPalindrome("ac"); str != "a" {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func longestPalindrome(s string) string {
	return ""
}
