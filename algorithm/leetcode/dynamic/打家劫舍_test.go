package dynamic

import (
	"testing"
)

func TestRob(t *testing.T) {
	if num := rob([]int{1, 2, 3, 1}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob([]int{2, 7, 9, 3, 1}); num != 12 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob([]int{1, 2, 3, 4, 5, 6, 4}); num != 13 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob([]int{2, 1, 1, 2}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	dp := make([][2]int, length)
	dp[0][1] = nums[0]
	dp[1][0] = dp[0][1]
	dp[1][1] = nums[1]
	for i := 2; i < length; i++ {
		dp[i][0] = max3(dp[i-1][1], dp[i-2][0])
		dp[i][1] = max3(dp[i-1][0], dp[i-2][0]) + nums[i]
	}
	return max3(dp[length-1][0], dp[length-1][1])
}
