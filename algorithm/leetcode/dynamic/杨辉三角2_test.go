package dynamic

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	arr := getRow(3)
	t.Log(arr)
	arr = getRow(1)
	t.Log(arr)
	arr = getRow(0)
	t.Log(arr)
	t.Log("测试通过")
}

func getRow(rowIndex int) []int {
	return getIndexRow(rowIndex + 1)
}

func getIndexRow(rowIndex int) []int {
	initRows(rowIndex)
	ret := make([]int, rowIndex)
	for i := 0; i < rowIndex; i++ {
		left := getPoint(rowIndex-1, i-1)
		right := getPoint(rowIndex-1, i)
		ret[i] = left + right
	}
	return ret
}

var rows [][]int

func getPoint(x, y int) int {
	if y == -1 {
		return 0
	}
	if x == 0 {
		return 1
	}
	if y == x {
		return 0
	}
	left := getValidPoint(x-1, y-1)
	right := getValidPoint(x-1, y)
	return left + right
}

func initRows(n int) {
	rows = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		row := make([]int, 0, n)
		for j := 0; j < n; j++ {
			row = append(row, -1)
		}
		rows = append(rows, row)
	}
}

func getValidPoint(x, y int) int {
	if x < 0 || y < 0 || y == len(rows[x]) {
		return 0
	}
	ret := rows[x][y]
	if ret == -1 {
		ret = getPoint(x, y)
		rows[x][y] = ret
	}
	return ret
}
