package array

import "testing"

func search(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		}
	}
	return count
}

func TestSearch(t *testing.T) {
	arr := []int{5, 7, 7, 8, 8, 10}
	n := search(arr, 8)
	if n != 2 {
		t.Fatal("测试失败")
		return
	}
	arr = []int{5, 7, 7, 8, 8, 10}
	n = search(arr, 6)
	if n != 0 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
