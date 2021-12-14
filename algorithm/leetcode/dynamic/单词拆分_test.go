package dynamic

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	if ok := wordBreak("leetcode", []string{"leet","code"}); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := wordBreak("applepenapple", []string{"apple","pen"}); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

func wordBreak(s string, wordDict []string) bool {
	return false
}

func findValid(index int,s string, wordDict []string) []int {
	length := len(s)
	ret := make([]int,0)
	for _,dict := range wordDict {
		dictLength := len(dict)
		if index + dictLength > length {
			continue
		}
		for i := 0; i < dictLength;i++ {
			if s[index+i] != dict[dictLength] {
				continue
			}
		}
		ret = append(ret,len(dict))
	}
	return ret
}
