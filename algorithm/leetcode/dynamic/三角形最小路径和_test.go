package dynamic

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	if num := minimumTotal([][]int{{1}, {2, 3}, {4, 5, 6}}); num != 7 {
		t.Fatal("测试不通过")
		return
	}
	if num := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}); num != 11 {
		t.Fatal("测试不通过")
		return
	}
	if num := minimumTotal([][]int{{-10}}); num != -10 {
		t.Fatal("测试不通过")
		return
	}
	if num := minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}}); num != -1 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1

f(0,0)=min(f(1,0),f(1,1))
*/
func minimumTotal(triangle [][]int) int {
	initMinimumTotalSaver(len(triangle))
	return getMinimumTotal(0, 0, triangle)
}

var minimumTotalSaver [][]int

func getMinimumTotal(x, y int, triangle [][]int) int {
	if x == len(triangle) {
		return 0
	}
	if len(triangle[x]) == y {
		return 10000
	}
	left := getMinimumTotalSaver(x+1, y, triangle) + triangle[x][y]
	right := getMinimumTotalSaver(x+1, y+1, triangle) + triangle[x][y]
	if left < right {
		return left
	}
	return right
}

func initMinimumTotalSaver(n int) {
	minimumTotalSaver = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		ret := make([]int, 0, n)
		for j := 0; j < n; j++ {
			ret = append(ret, -10001)
		}
		minimumTotalSaver = append(minimumTotalSaver, ret)
	}
}

func getMinimumTotalSaver(x, y int, triangle [][]int) int {
	if x == len(triangle) {
		return 0
	}
	if len(triangle[x]) == y {
		return 10000
	}
	num := minimumTotalSaver[x][y]
	if num == -10001 {
		num = getMinimumTotal(x, y, triangle)
		minimumTotalSaver[x][y] = num
	}
	return num
}
