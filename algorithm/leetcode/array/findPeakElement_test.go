package array

import "testing"

func FindPeakElement(nums []int) int {
	length := len(nums)
	if length == 1 || length == 2 {
		return -1
	}
	mid := length / 2
	if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		return mid
	}
	num := findPeakElement(nums, 0, mid-1)
	if num != -1 {
		return num
	}
	num = findPeakElement(nums, mid+1, length-1)
	if num != -1 {
		return num
	}
	return -1
}

func findPeakElement(nums []int, left, right int) int {
	length := right - left
	if length == 1 || length == 2 {
		return -1
	}
	mid := length / 2
	if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
		return mid
	}
	num := findPeakElement(nums, 0, mid-1)
	if num != -1 {
		return num
	}
	num = findPeakElement(nums, mid+1, length-1)
	if num != -1 {
		return num
	}
	return -1
}

func TestFindPeakElement(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	num := FindPeakElement(nums)
	if num != 2 {
		t.Fatal("查找失败:", num, "正确数组应该为2")
	}
	nums = []int{1, 2, 1, 3, 5, 6, 4}
	if num != 1 && num != 5 {
		t.Fatal("查找失败:", num, "正确数组应该为1或者5")
	}
	t.Log("测试成功")
}
