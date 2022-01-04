package dynamic

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	if num := longestCommonSubsequence("abcde", "ace"); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := longestCommonSubsequence("abc", "abc"); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := longestCommonSubsequence("abc", "def"); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := longestCommonSubsequence("bl", "yby"); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	if num := longestCommonSubsequence("bsbininm", "jmjkbkjkv"); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

var longestCommonSubsequenceSaver [][]int
var textOne string
var textTwo string

func longestCommonSubsequence(text1 string, text2 string) int {
	initLongestCommonSubsequenceSaver(len(text1), len(text2))
	textOne = text1
	textTwo = text2
	num := move(0, 0)
	return num
}

func initLongestCommonSubsequenceSaver(length1, length2 int) {
	longestCommonSubsequenceSaver = make([][]int, 0, length1+1)
	for i := 0; i < length1+1; i++ {
		arr := make([]int, 0, length2+1)
		for j := 0; j < length2+1; j++ {
			arr = append(arr, -1)
		}
		longestCommonSubsequenceSaver = append(longestCommonSubsequenceSaver, arr)
	}
}

func longestCommon(index1, index2 int, text1, text2 string) int {
	if index1 == len(text1) || index2 == len(text2) {
		return 0
	}
	a1 := longestCommonSubsequenceSaver[index1+1][index2]
	if a1 == -1 {
		a1 = longestCommon(index1+1, index2, text1, text2)
		if index1+1 < len(text1) {
			if text1[index1+1] == text2[index2] {
				a1 += 1
			}
		}
		longestCommonSubsequenceSaver[index1+1][index2] = a1
	}
	a2 := longestCommonSubsequenceSaver[index1][index2+1]
	if a2 == -1 {
		a2 = longestCommon(index1, index2+1, text1, text2)
		if index2+1 < len(text2) {
			if text1[index1] == text2[index2+1] {
				a2 += 1
			}
		}
		longestCommonSubsequenceSaver[index1][index2+1] = a2
	}
	return max(a1, a2)
}

func TestName(t *testing.T) {
	s1 := "absbd"
	s2 := "b"
	initLongestCommonSubsequenceSaver(len(s1), len(s2))
	num := move(0, 0)
	t.Log(num)
}

func move(index1, index2 int) int {
	if index1 == len(textOne) || index2 == len(textTwo) {
		return 0
	}
	if textOne[index1] == textTwo[index2] {
		return move(index1+1, index2+1) + 1
	}
	a1 := longestCommonSubsequenceSaver[index1+1][index2]
	if a1 == -1 {
		a1 = move(index1+1, index2)
		longestCommonSubsequenceSaver[index1+1][index2] = a1
	}
	a2 := longestCommonSubsequenceSaver[index1][index2+1]
	if a2 == -1 {
		a2 = move(index1, index2+1)
		longestCommonSubsequenceSaver[index1][index2+1] = a2
	}
	return max(a1, a2)
}
