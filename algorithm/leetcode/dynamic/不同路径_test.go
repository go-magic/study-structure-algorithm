package dynamic

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	if num := uniquePaths(3, 7); num != 28 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths(3, 2); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths(7, 3); num != 28 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths(3, 3); num != 6 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

f(0,0) = f(0,1) + f(1,0)
*/

var length int
var hight int
var allPathsSaver [][]int

func uniquePaths(m int, n int) int {
	length = m
	hight = n
	initSaver(m, n)
	return allPaths(0, 0)
}

func initSaver(x, y int) {
	allPathsSaver = make([][]int, 0, x)
	for i := 0; i < x; i++ {
		saver := make([]int, 0, y)
		for j := 0; j < y; j++ {
			saver = append(saver, -1)
		}
		allPathsSaver = append(allPathsSaver, saver)
	}
}

/*
保存中间变量
*/
func allPaths(x, y int) int {
	if x == length || y == hight {
		return 0
	}
	if x == length-1 && y == hight-1 {
		return 1
	}
	right := getLocaltion(x+1, y)
	boom := getLocaltion(x, y+1)
	return right + boom
}

func getLocaltion(x, y int) int {
	if x == length || y == hight {
		return 0
	}
	saver := allPathsSaver[x][y]
	if saver == -1 {
		saver = allPaths(x, y)
		allPathsSaver[x][y] = saver
	}
	return saver
}
