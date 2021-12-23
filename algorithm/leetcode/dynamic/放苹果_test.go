package dynamic

import (
	"testing"
)

func TestApple(t *testing.T) {
	arr := setApple(2,2)
	t.Log(arr)
}

var (
	appleNum int
	diskNum int
)

func setApple(m,n int) int {
	createDP(m,n)
	ans := 0
	for i := 1; i <= n; i++ {
		ans += solve(m, i)
	}
	return ans
}

func set(at,total int) int {
	if at == diskNum-1 {
		return 1
	}
	if total < 1 {
		return 1
	}
	nums := 0
	for i := 0; i < total+1;i++ {
		num := set(at+1,total-i)
		nums += num
	}
	return nums
}

var dp [][]int

func createDP(m,n int){
	dp = make([][]int,0,m+1)
	for i := 0;i < m+1;i++ {
		arr := make([]int,n+1)
		for j := 0;j < n+1;j++{
			arr[j] = -1
		}
		dp = append(dp,arr)
	}
}

func solve(m,n int) int {
	if dp[m][n] != -1 {
		return dp[m][n]
	}
	if m < n {
		dp[m][n] = 0
		return dp[m][n]
	}
	if m == n || n == 1 {
		dp[m][n] = 1
		return dp[m][n]
	}
	dp[m][n] = 0
	for i := 1; i <= n; i++ {
		dp[m][n] += solve(m - n, i)
	}
	return dp[m][n]
}
