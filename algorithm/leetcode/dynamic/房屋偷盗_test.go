package dynamic

import (
	"testing"
)

func TestCommonRob(t *testing.T) {
	if num := commonRob([]int{1, 2, 3, 1}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := commonRob([]int{2, 7, 9, 3, 1}); num != 12 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
解题思路
子问题:

当前不盗则上一次可以盗，也可以不盗，选二者最大值。dp[i][0] = max(dp[i-1][0], dp[i-1][1])。
当前盗则上一次不可以盗，但是需要加上这一次盗的金额。dp[i][1] = dp[i-1][0] + nums[i]。

*/
func commonRob(nums []int) int {
	length := len(nums)
	dp := make([][2]int, length)
	dp[0][1] = nums[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[length-1][0], dp[length-1][1])
}
