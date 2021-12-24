package dynamic

import (
	"testing"
)

func TestUniquePaths1(t *testing.T) {
	if num := uniquePaths2(3, 7); num != 28 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths2(7, 3); num != 28 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths2(3, 3); num != 6 {
		t.Fatal("测试不通过")
		return
	}
	if num := uniquePaths2(3, 2); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

var (
	dpPath [][]int
	xPath  int
	yPath  int
)

func uniquePaths1(m int, n int) int {
	xPath = m
	yPath = n
	initPaths(m, n)
	return step(0, 0)
}

func initPaths(m, n int) {
	dpPath = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dpPath[i] = make([]int, n+1)
	}
}

func step(x, y int) int {
	if x == xPath || y == yPath {
		return 0
	}
	if x == xPath-1 && y == yPath-1 {
		return 1
	}
	right := dpPath[x+1][y]
	if right == 0 {
		right = step(x+1, y)
		dpPath[x+1][y] = right
	}
	below := dpPath[x][y+1]
	if below == 0 {
		below = step(x, y+1)
		dpPath[x][y+1] = below
	}
	return right + below
}

//官方(非递归)
func uniquePaths2(m, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[0][j] = 1
			}
			if j == 0 {
				dp[i][0] = 1
			}

			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
