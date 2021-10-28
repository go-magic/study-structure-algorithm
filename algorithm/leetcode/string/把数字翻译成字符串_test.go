package string

import (
	"testing"
)

func TestTranslateNum(t *testing.T) {
	if num := translateNum(12258); num != 5 {
		t.Fatal("测试失败")
		return
	}
	if num := translateNum(25); num != 2 {
		t.Fatal("测试失败")
		return
	}
	if num := translateNum(506); num != 1 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	arr := getArr(num)
	ret := find(arr[0], 1, arr)
	if arr[0] != 0 {
		ret += find(arr[0]*10+arr[1], 2, arr)
	}
	return ret
}

/*
506
5 + (06)
5 + 0 + 6
*/
func find(mod, i int, arr []int) int {
	if mod > 25 {
		return 0
	}
	if i >= len(arr)-1 {
		return 1
	}
	ret := find(arr[i], i+1, arr)
	if arr[i] != 0 {
		ret += find(arr[i]*10+arr[i+1], i+2, arr)
	}
	return ret
}

func getArr(num int) []int {
	ret := make([]int, 0)
	for num != 0 {
		ret = append(ret, num%10)
		num = num / 10
	}
	return reverse(ret)
}

func reverse(arr []int) []int {
	length := len(arr)
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[length-1-i] = arr[i]
	}
	return ret
}
