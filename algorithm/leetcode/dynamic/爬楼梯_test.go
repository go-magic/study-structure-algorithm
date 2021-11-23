package dynamic

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	if num := climbStairs(1); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := climbStairs(2); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	if num := climbStairs(3); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := climbStairs(4); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
f(n) = f(n-1) + f(n-2)
*/

var climbStairsSaver []int

func climbStairs(n int) int {
	climbStairsSaver = make([]int, n+1)
	return getClimbStairs(n)
}

func getClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	one := climbStairsSaver[n-1]
	if one == 0 {
		one = getClimbStairs(n - 1)
		climbStairsSaver[n-1] = one
	}
	two := climbStairsSaver[n-2]
	if two == 0 {
		two = getClimbStairs(n - 2)
		climbStairsSaver[n-2] = two
	}
	return one + two
}
