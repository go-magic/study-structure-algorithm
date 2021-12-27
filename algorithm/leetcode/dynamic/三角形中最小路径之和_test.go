package dynamic

import (
	"testing"
)

func TestMinimumTotals(t *testing.T) {
	if num := minimumTotals([][]int{{2}, {3,4}, {6,5,7},{4,1,8,3}}); num != 11 {
		t.Fatal("测试不通过")
		return
	}
	if num := minimumTotals([][]int{{-10}}); num != -10 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
   2
  3 4
 6 5 7
4 1 8 3
*/

var minimumTotalsSaver [][]int
func minimumTotals(triangle [][]int) int {
	initMinimumTotalsSaver(triangle)
	return minimumTotalsSolve(0,0,triangle)
}

func initMinimumTotalsSaver(triangle [][]int)  {
	length := len(triangle)
	minimumTotalsSaver = make([][]int,0,length+1)
	for i := 0 ;i < length+1;i++ {
		saver := make([]int,0,length+1)
		for j := 0; j < length+1;j++ {
			saver = append(saver,-1)
		}
		minimumTotalsSaver = append(minimumTotalsSaver,saver)
	}
}

func minimumTotalsSolve(x,y int,triangle [][]int) int {
	length := len(triangle)
	if y == len(triangle) {
		return 0
	}
	if x == len(triangle[length-1]) {
		return 0
	}
	left := minimumTotalsSaver[x][y+1]
	if left == -1 {
		left = minimumTotalsSolve(x,y+1,triangle)
		minimumTotalsSaver[x][y+1] = left
	}
	right := minimumTotalsSaver[x+1][y+1]
	if right == - 1 {
		right = minimumTotalsSolve(x+1,y+1,triangle)
		minimumTotalsSaver[x+1][y+1] = right
	}
	return min(left,right) + triangle[y][x]
}