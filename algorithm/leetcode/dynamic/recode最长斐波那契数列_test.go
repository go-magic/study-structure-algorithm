package dynamic

import (
	"testing"
)

func TestLenLongestFibSubSeq(t *testing.T) {
	if num := lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	if num := lenLongestFibSubseq([]int{1, 3, 7, 10, 11, 12, 14, 17, 18}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := lenLongestFibSubseq([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}); num != 5 {
		t.Fatal("测试不通过:", num)
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
	saver := make(map[int]int)
	for i, val := range arr {
		saver[val] = i
	}
	length := len(arr)
	maxLength := 0
	for k := 0; k < length; k++ {
		count := 0
		for i := k; i < length; i++ {
			for j := i; j < length; j++ {
				index, exist := saver[arr[i]+arr[j]]
				if !exist {
					break
				}
				i = j
				j = index
				count++
			}
		}
		if maxLength < count {
			maxLength = count
		}
	}
	if maxLength != 0 {
		maxLength += 2
	}
	return maxLength
}

/*
f(arr[i])=dp[i-1]+dp[i-2]
*/
func selectIndex(index1, index2 int, ok bool, arr []int, saver map[int]int) int {
	if index2 == len(arr) {
		return 0
	}
	val := arr[index1] + arr[index2]
	index, exist := saver[val]
	if !exist {
		if ok {
			return 0
		}
		return selectIndex(index1, index2+1, false, arr, saver)
	}
	a1 := selectIndex(index2, index, true, arr, saver) + 1
	a2 := 0
	if !ok {
		a2 = selectIndex(index1, index2+1, ok, arr, saver)
	}
	return max(a1, a2)
}

var indexSaver [][]int

func selectIndex1(arr []int) int {
	length := len(arr)
	saver := make(map[int]int)
	for i, val := range arr {
		saver[val] = i
	}
	maxLength := 0
	for i := 0; i < length; i++ {
		a := selectIndex(i, i+1, false, arr, saver)
		if maxLength < a {
			maxLength = a
		}
	}
	if maxLength != 0 {
		maxLength = maxLength + 2
	}
	return maxLength
}

func initIndexSaver(length int) {
	indexSaver = make([][]int, 0, length+1)
	for i := 0; i < length+1; i++ {
		arr := make([]int, 0, length+1)
		for j := 0; j < length+1; j++ {
			arr = append(arr, -1)
		}
		indexSaver = append(indexSaver, arr)
	}
}

func compare(s1, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)
	if length1 != length2 {
		return false
	}
	m := 0
	n := length1 - 1
	for i := 0; i < length1; i++ {
		if s1[m] != s2[n] {
			return false
		}
		m++
		n--
	}
	return true
}

func TestCompare(t *testing.T) {
	t.Log(compare("as1sd", "ds1sw"))
}

func TestSelectIndex(t *testing.T) {
	saver := make(map[int]int)
	for i, val := range arr {
		saver[val] = i
	}
	if num := selectIndex1([]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 12, 17, 19, 27, 29}); num != 5 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := selectIndex1([]int{1, 2, 3, 4, 5, 6, 7, 8}); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	if num := selectIndex1([]int{1, 3, 7, 10, 11, 12, 14, 17, 18}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := selectIndex1([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}); num != 5 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := selectIndex1([]int{2, 4, 5, 6, 7, 8, 11, 13, 14, 15, 21, 22, 34}); num != 5 {
		t.Fatal("测试不通过:", num)
		return
	}
	t.Log("测试通过")
}
