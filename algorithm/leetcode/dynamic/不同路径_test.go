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

func uniquePaths(m int, n int) int {
	length = m
	hight = n
	allPathsSaver = make(map[int]map[int]int)
	return allPaths2(0, 0)
}

func allPaths1(x, y int) int {
	if x == length || y == hight {
		return 0
	}
	if x == length-1 && y == hight-1 {
		return 1
	}
	right := allPaths(x+1, y)
	boom := allPaths(x, y+1)
	return right + boom
}

/*
保存中间变量
*/
func allPaths2(x, y int) int {
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
	var now int
	_, ok := allPathsSaver[x]
	if !ok {
		allPathsSaver[x] = make(map[int]int)
		now = allPaths(x, y)
		allPathsSaver[x][y] = now
	}
	if ok {
		now, ok = allPathsSaver[x][y]
		if !ok {
			now = allPaths(x, y)
			allPathsSaver[x][y] = now
		}
	}
	return now
}

func allPaths(x, y int) int {
	if x == length || y == hight {
		return 0
	}
	if x == length-1 && y == hight-1 {
		return 1
	}
	var (
		right int
		boom  int
	)
	_, ok := allPathsSaver[x+1]
	if !ok {
		allPathsSaver[x+1] = make(map[int]int)
		right = allPaths(x+1, y)
		allPathsSaver[x+1][y] = right
	}
	if ok {
		right, ok = allPathsSaver[x+1][y]
		if !ok {
			right = allPaths(x+1, y)
			allPathsSaver[x+1][y] = right
		}
	}
	_, ok = allPathsSaver[x]
	if !ok {
		allPathsSaver[x] = make(map[int]int)
		boom = allPaths(x, y+1)
		allPathsSaver[x][y+1] = boom
	}
	if ok {
		boom, ok = allPathsSaver[x][y+1]
		if !ok {
			boom = allPaths(x, y+1)
			allPathsSaver[x][y+1] = boom
		}
	}
	return right + boom
}

var allPathsSaver map[int]map[int]int
