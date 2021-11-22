package dynamic

import (
	"testing"
)

func TestJump1(t *testing.T) {
	//可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
	if ok := canJump([]int{2, 3, 1, 1, 4}); !ok {
		t.Fatal("测试不通过")
		return
	}
	//无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
	if ok := canJump([]int{3, 2, 1, 0, 4}); ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := canJump([]int{2, 3, 0, 1, 0, 1}); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。
*/
func canJump(nums []int) bool {
	jumpOutSaver = make([]int, 30001)
	return canJumpOut_2(0, nums)
}

/*
第一步每个位置有当前这么多种走法(将问题分解成子问题,f(0)=max(f(1),f(2),f(3)))
*/
func canJumpOut_1(index int, nums []int) bool {
	if index >= len(nums)-1 {
		return true
	}
	num := nums[index]
	//每个位置有当前这么多种走法
	for i := 1; i <= num; i++ {
		if ok := canJumpOut(index+i, nums); ok {
			return true
		}
	}
	return false
}

/*
增加保存中间相同结果数据
*/

var jumpOutSaver []int

func canJumpOut_2(index int, nums []int) bool {
	if index >= len(nums)-1 {
		return true
	}
	num := nums[index]
	//每个位置有当前这么多种走法
	for i := 1; i <= num; i++ {
		saver := jumpOutSaver[index+i]
		if saver == 1 {
			return true
		}
		if saver == 2 {
			continue
		}
		if ok := canJumpOut_2(index+i, nums); ok {
			jumpOutSaver[index+i] = 1
			return true
		}
		//走过且不通过的设置为2
		jumpOutSaver[index+i] = 2
	}
	return false
}

func canJumpOut(index int, nums []int) bool {
	if index >= len(nums)-1 {
		return true
	}
	num := nums[index]
	//每个位置有当前这么多种走法
	for i := 1; i <= num; i++ {
		if ok := canJumpOut(index+i, nums); ok {
			return true
		}
	}
	return false
}
