package dynamic

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {
	if ok := isInterleave("aabcc", "dbbca", "aadbbcbcac"); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("aabcc", "dbbca", "aadbbbaccc"); ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("", "", ""); !ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return false
}
