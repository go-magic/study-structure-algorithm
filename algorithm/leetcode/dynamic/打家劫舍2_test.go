package dynamic

import "testing"

func TestRob2(t *testing.T) {
	if num := rob2([]int{1, 2, 3, 1}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob2([]int{2, 7, 9, 3, 1}); num != 12 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob2([]int{1, 2, 3, 4, 5, 6, 4}); num != 13 {
		t.Fatal("测试不通过")
		return
	}
	if num := rob2([]int{2, 1, 1, 2}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额

*/
func rob2(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	dp := make([][2]int,length)
	dp[0][1] = nums[1]
	for i := 2; i < length;i++ {
		dp[i][0] = max(dp[i-1][1],dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return  max(dp[length-1][1],dp[length-1][0])
}