package dynamic

import (
	"fmt"
	"testing"
)

func TestMinPathSums(t *testing.T) {
	if num := minPathSums([][]int{{1,3,1}, {1,5,1}, {4,2,1}}); num != 7 {
		t.Fatal("测试不通过")
		return
	}
	if num := minPathSums([][]int{{1,2,3},{4,5,6}}); num != 12 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
1,3,1
1,5,1
4,2,1
*/

/*
1,2,3
4,5,6
*/
func minPathSums(grid [][]int) int {
	length := len(grid)
	hight := len(grid[0])
	dp := make([][]int,length)
	for i := 0; i < length;i++ {
		dp[i] = make([]int,hight)
	}
	dp[length-1][hight-1] = grid[length-1][hight-1]
	for i := length-1; i >=0;i-- {
		for j := hight-1; j >= 0;j-- {
			if i == length-1 && j == hight-1 {
				continue
			}
			if i == length-1 {
				dp[i][j] = dp[i][j+1] + grid[i][j]
				continue
			}
			if j == hight-1 {
				dp[i][j] = dp[i+1][j] + grid[i][j]
				continue
			}
			dp[i][j] = min(dp[i+1][j],dp[i][j+1]) + grid[i][j]
		}
	}
	for _,v := range dp {
		fmt.Println(v)
	}
	return dp[0][0]
}
