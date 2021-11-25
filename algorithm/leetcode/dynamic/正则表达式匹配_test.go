package dynamic

import (
	"testing"
)

func TestDo(t *testing.T) {
	if ok := isMatch("ab", ".*c"); ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aaa", "ab*a*c*a"); !ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aa", "a"); ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aa", "a*"); !ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aa", ".a"); !ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aa", ".*"); !ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aac", ".*"); ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("aab", "c*a*b"); !ok {
		t.Fatal("测试失败")
		return
	}
	if ok := isMatch("mississippi", "mis*is*p*."); ok {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

/*

 */
func isMatch(s string, p string) bool {
	return getMatch(0, 0, s, p, "")
}

func getMatch(index1, index2 int, s string, p string, star string) bool {
	if index1 == len(s) {
		if index2 == len(p) {
			return true
		}
		return false
	}
	if index2 == len(p) {
		return false
	}
	if star != "" {
		if string(s[index1]) == star {
			if ok := getMatch(index1+1, index2, s, p, star); ok {
				return true
			}
		}
		return false
	}
	c := p[index2]
	if c == '.' {
		return getMatch(index1+1, index2+1, s, p, "")
	}
	if c == '*' {
		if ok := getMatch(index1+1, index2, s, p, string(s[index1])); ok {
			return true
		}
		if ok := getMatch(index1+1, index2+1, s, p, ""); ok {
			return true
		}
		return false
	}
	if c == s[index1] {
		if index2+1 < len(p) && p[index2+1] == '*' {
			if ok := getMatch(index1, index2+2, s, p, ""); ok {
				return true
			}
		}
		return getMatch(index1+1, index2+1, s, p, "")
	}
	if index2+1 < len(p) && p[index2+1] == '*' {
		return getMatch(index1, index2+2, s, p, "")
	}
	return false
}
