package dynamic

import (
	"testing"
)

func TestLenLongestFibSubSeq(t *testing.T) {
	if num := lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	if num := lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
如果序列X_1, X_2, ..., X_n满足下列条件，就说它是斐波那契式的：

n >= 3
对于所有i + 2 <= n，都有X_i + X_{i+1} = X_{i+2}
给定一个严格递增的正整数数组形成序列 arr，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回0 。

（回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如，[3, 5, 8][3, 4, 5, 6, 7, 8]的一个子序列）
*/
func lenLongestFibSubseq(arr []int) int {
	return 0
}
