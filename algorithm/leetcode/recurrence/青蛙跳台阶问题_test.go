package recurrence

import "testing"

func TestNumWays(t *testing.T) {
	if n := numWays(0); n != 1 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays(1); n != 1 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays(2); n != 2 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays(3); n != 3 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays(4); n != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
func TestNumWays1(t *testing.T) {
	if n := numWays1(0); n != 1 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays1(1); n != 1 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays1(2); n != 2 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays1(3); n != 3 {
		t.Fatal("测试失败")
		return
	}
	if n := numWays1(4); n != 5 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1 := 1
	n2 := 2
	var ret int
	for i := 2; i < n; i++ {
		ret = n1 + n2
		ret = ret % 1000000007
		n1 = n2
		n2 = ret
	}
	return ret
}

func numWays1(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return numWays1(n-1) + numWays1(n-2)
}
