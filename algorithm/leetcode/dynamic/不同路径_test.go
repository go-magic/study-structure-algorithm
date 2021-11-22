package dynamic

import (
	"testing"
)

func TestJump1(t *testing.T) {
	//可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
	if ok := canJump([]int{2, 3, 1, 1, 4}); !ok {
		t.Fatal("测试不通过")
		return
	}
	//无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
	if ok := canJump([]int{3, 2, 1, 0, 4}); ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := canJump([]int{2, 3, 0, 1, 0, 1}); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}
