package dynamic

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	if num := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{0, 0}, {0, 1}}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{0}}); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{1}}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{0, 1}}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{0}, {1}}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePathsWithObstacles([][]int{{1, 0}}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

思考: f(0,0) = f(0,1) + f(1,0)

*/
var obstaclesLength int
var obstaclesHight int
var obstaclesSaver [][]int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	obstaclesLength = len(obstacleGrid)
	obstaclesHight = len(obstacleGrid[0])
	if obstaclesLength == 1 && obstaclesHight == 1 && obstacleGrid[0][0] == 0 {
		return 1
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	createObstaclesSaver(obstaclesLength, obstaclesHight)
	rigth := getPaths(0, 1, obstacleGrid)
	boom := getPaths(1, 0, obstacleGrid)
	return rigth + boom
}

func createObstaclesSaver(x, y int) {
	obstaclesSaver = make([][]int, 0, x)
	for i := 0; i < x; i++ {
		saver := make([]int, 0, y)
		for j := 0; j < y; j++ {
			saver = append(saver, -1)
		}
		obstaclesSaver = append(obstaclesSaver, saver)
	}
}

func getPaths(x, y int, obstacleGrid [][]int) int {
	if x == obstaclesLength || y == obstaclesHight {
		return 0
	}
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	if x == obstaclesLength-1 && y == obstaclesHight-1 {
		return 1
	}
	right := getPathsWithArrays(x+1, y, obstacleGrid)
	boom := getPathsWithArrays(x, y+1, obstacleGrid)
	return right + boom
}

func getPathsWithArrays(x, y int, obstacleGrid [][]int) int {
	if x == obstaclesLength || y == obstaclesHight {
		return 0
	}
	saver := obstaclesSaver[x][y]
	if saver == -1 {
		saver = getPaths(x, y, obstacleGrid)
		obstaclesSaver[x][y] = saver
		return saver
	}
	return saver
}
