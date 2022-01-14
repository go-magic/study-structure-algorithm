package dynamic

import "testing"

func TestCuttingRope1(t *testing.T) {
	if num := cuttingRope1(2); num != 1 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope1(3); num != 2 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope1(4); num != 4 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope1(10); num != 36 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope1(58); num != 953271190 {
		t.Fatal(num)
		return
	}
	t.Log("测试成功")
}

var cutSaver1 []int

func cuttingRope1(n int) int {
	cutSaver1 = make([]int, 0, n+1)
	for i := 0; i < n+1; i++ {
		cutSaver1 = append(cutSaver1, -1)
	}
	return subCuttingRope1(n)
}

func subCuttingRope1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	var maxValue = 1
	for i := 1; i < n; i++ {
		value := cutSaver1[n-i]
		if value == -1 {
			value = subCuttingRope1(n - i)
			cutSaver1[n-i] = value
		}
		value = value * i
		maxValue = max(max(maxValue, value), i*(n-i))
	}
	return maxValue
}
