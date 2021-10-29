package array

import "testing"

func TestMaxSubArray(t *testing.T) {
	if num := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); num != 6 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

var max int

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
}

func findMaxSubArray(start, end, value int, nums []int) int {
	if end == len(nums)-1 {
		return end - start, value
	}

}
