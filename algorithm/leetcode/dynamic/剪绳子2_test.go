package dynamic

import "testing"

func TestCuttingRope(t *testing.T) {
	if num := cuttingRope(2); num != 1 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope(3); num != 2 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope(4); num != 4 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope(10); num != 36 {
		t.Fatal(num)
		return
	}
	if num := cuttingRope(120); num != 953271190 {
		t.Fatal(num)
		return
	}
	t.Log("测试成功")
}

var cutSaver []int64

func cuttingRope(n int) int {
	cutSaver = make([]int64, 0, n+1)
	for i := 0; i < n+1; i++ {
		cutSaver = append(cutSaver, -1)
	}
	var maxValue int64 = 1
	for i := 1; i < n; i++ {
		value := subCuttingRope(n-i) % 1000000007
		value = (value * int64(i)) % 1000000007
		maxValue = max64(max64(maxValue, value), int64(i*(n-i))) % 1000000007
	}
	return int(maxValue) % 1000000007
}

func subCuttingRope(n int) int64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	var maxValue int64 = 1
	for i := 1; i < n; i++ {
		value := cutSaver[n-i]
		if value == -1 {
			value = subCuttingRope(n - i)
			cutSaver[n-i] = value
		}
		value = value * int64(i)
		maxValue = max64(max64(maxValue, value), int64(i*(n-i)))
	}
	return maxValue
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
