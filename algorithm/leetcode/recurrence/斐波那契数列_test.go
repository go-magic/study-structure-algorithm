package recurrence

import "testing"

func TestFib(t *testing.T) {
	if n := fib(2); n != 1 {
		t.Fatal("测试失败")
		return
	}
	if n := fib(5); n != 5 {
		t.Fatal("测试失败")
		return
	}
	if n := fib(6); n != 8 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

func fib(n int) int {
	return mod(n) % 1000000007
}

func mod(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1 := 0
	n2 := 1
	for i := 2; i < n; i++ {
		ret := n1 + n2
		ret = ret % 1000000007
		n1 = n2
		n2 = ret
	}
	return n1 + n2
}
