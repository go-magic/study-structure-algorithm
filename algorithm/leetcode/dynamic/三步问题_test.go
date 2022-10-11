package dynamic

import "testing"

/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。

示例1:

 输入：n = 3
 输出：4
 说明: 有四种走法
示例2:

 输入：n = 5
 输出：13
提示:

n范围在[1, 1000000]之间

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/three-steps-problem-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func waysToStep(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return (waysToStep(n-1) + waysToStep(n-2) + waysToStep(n-3)) % 1000000007
}

func waysToStep1(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}
	x := 1
	y := 2
	z := 4
	for i := 4; i < n; i++ {
		tmp := x
		x = y
		y = z
		z = (tmp + x + y) % 1000000007
	}
	return (x + y + z) % 1000000007
}

func TestWaysToStep(t *testing.T) {
	if length := waysToStep1(1); length != 1 {
		t.Fatal("测试失败")
		return
	}
	if length := waysToStep1(2); length != 2 {
		t.Fatal("测试失败")
		return
	}
	if length := waysToStep1(3); length != 4 {
		t.Fatal("测试失败")
		return
	}
	if length := waysToStep1(4); length != 7 {
		t.Fatal("测试失败")
		return
	}
	if length := waysToStep1(5); length != 13 {
		t.Fatal("测试失败")
		return
	}
	if length := waysToStep1(61); length != 752119970 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
