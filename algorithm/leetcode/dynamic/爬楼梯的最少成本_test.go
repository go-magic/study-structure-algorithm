package dynamic

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	if num := minCostClimbingStairs([]int{10, 15, 20}); num != 15 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}); num != 6 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := minCostClimbingStairs([]int{1, 100}); num != 1 {
		t.Fatal("测试不通过:", num)
		return
	}
	t.Log("测试通过")
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	if length == 1 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < length; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[length-2], dp[length-1])
}

func minCostClimbingStairsStep(index int, cost []int) int {
	if index >= len(cost) {
		return 0
	}
	one := minCostClimbingStairsStep(index+1, cost) + cost[index]
	two := minCostClimbingStairsStep(index+2, cost) + cost[index]
	return min(one, two)
}
