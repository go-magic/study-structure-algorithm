package dynamic

import (
	"testing"
)

func TestMaxValue(t *testing.T) {
	/*
		1, 3, 1
	*/
	if value := maxGiftValue([][]int{{1, 3, 1}}); value != 5 {
		t.Fatal(value)
		return
	}
	/*
		1, 3, 1
		1, 5, 1
	*/
	if value := maxGiftValue([][]int{{1, 3, 1}, {1, 5, 1}}); value != 10 {
		t.Fatal(value)
		return
	}
	/*
		1, 3, 1
		1, 5, 1
		4, 2, 1
	*/
	if value := maxGiftValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}); value != 12 {
		t.Fatal(value)
		return
	}
	/*
		1, 2
		1, 1
	*/
	if value := maxGiftValue([][]int{{1, 2}, {1, 1}}); value != 4 {
		t.Fatal(value)
		return
	}
	t.Log("测试通过")
}

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

f(0,0) = max(f(0,1)+grid[0][1],f(1,0)+grid[0][1])

*/
func maxGiftValue(grid [][]int) int {
	y := len(grid)
	if y == 0 {
		return 0
	}
	savers = make(map[int]map[int]int)
	rigth := giftValue(1, 0, grid) + grid[0][0]
	boom := giftValue(0, 1, grid) + grid[0][0]
	if rigth > boom {
		return rigth
	}
	return boom
}

func giftValue(x, y int, grid [][]int) int {
	if x == len(grid) || y == len(grid[0]) {
		return 0
	}
	var (
		rigth int
		boom  int
		exist bool
	)
	rigth, exist = checkExist(x+1, y)
	if !exist {
		rigth = giftValue(x+1, y, grid)
		savers[x+1][y] = rigth
	}
	rigth += grid[x][y]
	boom, exist = checkExist(x, y+1)
	if !exist {
		boom = giftValue(x, y+1, grid)
		savers[x][y+1] = boom
	}
	boom += grid[x][y]
	if rigth > boom {
		return rigth
	}
	return boom
}

var savers = make(map[int]map[int]int)

func checkExist(x, y int) (int, bool) {
	saver, exist := savers[x]
	if !exist {
		savers[x] = make(map[int]int)
		return 0, exist
	}
	value, exist := saver[y]
	if !exist {
		return 0, exist
	}
	return value, true
}
