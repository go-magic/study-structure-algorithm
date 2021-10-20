package array

import "testing"

func missingNumber(nums []int) int {
	length := len(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return nums[length-1] + 1
}

func TestMissingNumber(t *testing.T) {
	arr := []int{0, 1, 3}
	n := missingNumber(arr)
	if n != 2 {
		t.Fatal("测试失败")
		return
	}
	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	n = missingNumber(arr)
	if n != 8 {
		t.Fatal("测试失败")
		return
	}
	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	n = missingNumber(arr)
	if n != 9 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
