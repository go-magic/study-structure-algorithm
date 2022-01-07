package dynamic

import (
	"testing"
)

func TestRob1(t *testing.T) {
	if num := rob1([]int{2, 1}); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob1([]int{2, 3, 2}); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob1([]int{2, 4, 2}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob1([]int{2, 4, 8, 1, 2, 1, 3, 1}); num != 15 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
子问题:分两种情况,
	1.第一个房子不刷，最后一个房子可刷可不刷。
	2.第一个房子刷，最后一个房子不刷。
求出两个dp的最大值，注意边界条件房子总数为1和2时。
*/
func rob1(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	one := rob3(nums)
	two := rob4(nums)
	return max(one, two)
}

func rob3(nums []int) int {
	length := len(nums)
	dp := make([][2]int, length)
	dp[1][0] = 0
	dp[1][1] = nums[1]
	for i := 2; i < length-1; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[length-2][0]+nums[length-1], dp[length-2][1])
}

func rob4(nums []int) int {
	length := len(nums)
	dp := make([][2]int, length)
	dp[1][0] = nums[0]
	dp[1][1] = nums[0]
	for i := 2; i < length-1; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	dp[length-1][0] = max(dp[length-2][0], dp[length-2][1])
	return max(dp[length-2][0], dp[length-2][1])
}
