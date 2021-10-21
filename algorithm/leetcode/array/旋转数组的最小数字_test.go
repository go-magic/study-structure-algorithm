package array

import "testing"

func TestMinArray(t *testing.T) {
	arr := []int{3, 4, 5, 1, 2}
	if a := minArray(arr); a != 1 {
		t.Fatal("测试失败")
		return
	}
	arr = []int{2, 2, 2, 0, 1}
	if a := minArray(arr); a != 0 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	temp := numbers[0]
	for _, val := range numbers {
		if temp > val {
			return val
		}
	}
	return temp
}
