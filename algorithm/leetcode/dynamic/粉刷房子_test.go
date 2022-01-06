package dynamic

import (
	"testing"
)

func TestMinCost(t *testing.T) {
	if num := minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19}}); num != 10 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := minCost([][]int{{7, 6, 2}}); num != 2 {
		t.Fatal("测试不通过:", num)
		return
	}
	t.Log("测试通过")
}

/*
解题思路
子问题，当前刷0则上一次可以刷1或2,取最小值。当前刷1则上一次可以刷0或2,取最小值。当前刷2则上一次可以刷0或1,则有:

f(n,0)=min(f(n-1,1),f(n-1,2)) + costs[n][0]
f(n,1)=min(f(n-1,0),f(n-1,2)) + costs[n][1]
f(n,2)=min(f(n-1,0),f(n-1,1)) + costs[n][2]

*/
func minCost(costs [][]int) int {
	length := len(costs)
	dp := make([][3]int, length)
	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]
	for i := 1; i < length; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	return min3(dp[length-1][0], dp[length-1][1], dp[length-1][2])
}
