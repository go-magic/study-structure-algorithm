package dynamic

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	arr := countBits(2)
	t.Log(arr)
	arr = countBits(3)
	t.Log(arr)
	arr = countBits(5)
	t.Log(arr)
	arr = countBits(6)
	t.Log(arr)
	arr = countBits(7)
	t.Log(arr)
}

func countBits(n int) []int {
	ret := make([]int, 0, n+1)
	dp := make([]int, n+1)
	ret = append(ret, 0)
	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			ret = append(ret, 1)
			dp[i] = 1
			continue
		}
		flag := 0
		if (i & 1) == 1 {
			flag = 1
		}
		temp := dp[i>>1] + flag
		ret = append(ret, temp)
		dp[i] = temp
	}
	return ret
}

/*
0,1,1,2,1,2,2,3,1
*/
func getBits(n int) int {
	return length
}

func TestBits(t *testing.T) {
	arr := countBits(0)
	t.Log(arr)
}
