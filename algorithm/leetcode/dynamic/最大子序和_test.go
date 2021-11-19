package dynamic

import (
	"testing"
)

func TestMaxSubArrays(t *testing.T) {
	if num := maxSubArrays([]int{-2, -1}); num != -1 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{2, -1, 2}); num != 3 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{1}); num != 1 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{0}); num != 0 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{-1}); num != -1 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{-100000}); num != -100000 {
		t.Fatal(num)
		return
	}
	if num := maxSubArrays([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); num != 6 {
		t.Fatal(num)
		return
	}
	t.Log("测试通过")
}

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/
func maxSubArrays(nums []int) int {
	length := len(nums)
	max := -100000
	temp := 0
	for i := 0; i < length; i++ {
		if nums[i] <= temp+nums[i] {
			temp += nums[i]
			if max < temp {
				max = temp
			}
		} else {
			temp = nums[i]
		}
	}
	if max < temp {
		max = temp
	}
	return max
}
