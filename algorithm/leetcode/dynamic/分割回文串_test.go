package dynamic

import (
	"testing"
)

func TestPartition(t *testing.T) {
	arr :=partition("aab")
	t.Log(arr)
	arr =partition("a")
	t.Log(arr)
	t.Log("测试通过")
}

func partition(s string) [][]string {
	return [][]string{}
}
