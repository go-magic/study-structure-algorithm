package array

import "testing"

func findRepeatNumber(nums []int) int {
	m := make([]int, 100001)
	for _, v := range nums {
		a := m[v]
		if a > 0 {
			return v
		}
		m[v]++
	}
	return -1
}

func TestFindRepeatNumber(t *testing.T) {
	arr := []int{3, 4, 2, 0, 0, 1}
	n := findRepeatNumber(arr)
	if n != 2 && n != 3 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func TestArr(t *testing.T) {
	arr := make([]int, 10)
	arr[0]++
	arr[1]++
	t.Log(arr)
}
