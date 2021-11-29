package dynamic

import (
"testing"
)

func TestMaxProfit3(t *testing.T) {
	if num := maxProfit2([]int{7,6,4,3,1}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit2([]int{1,2,3,4,5}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit2([]int{7,1,5,3,6,4}); num != 7 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit2([]int{2,1,2,0,1}); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
0表示卖,1表示买
*/
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max3(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max3(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}