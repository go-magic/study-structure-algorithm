package dynamic

import "testing"

/*
给定一个整数数组，找出总和最大的连续数列，并返回总和。

示例：

输入： [-2,1,-3,4,-1,2,1,-5,4]
输出： 6
解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶：

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/contiguous-sequence-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxSubArray1(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	x := 0
	y := nums[0]
	for i := 1; i < len(nums); i++ {
		x = max(x, y)
		y = max(y+nums[i], nums[i])
	}
	return max(x, y)
}

func TestMaxSubArray1(t *testing.T) {
	if length := maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); length != 6 {
		t.Fatal("测试失败")
		return
	}
	if length := maxSubArray1([]int{-2, 1, 3}); length != 4 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试成功")
}
