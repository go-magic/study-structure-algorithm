package dynamic

import "testing"

func TestFindSquare(t *testing.T) {
	t.Log(findSquare([][]int{
		{1, 0, 1},
		{0, 0, 1},
		{0, 0, 1},
	}))
	t.Log(findSquare([][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}))
	t.Log(findSquare([][]int{
		{1},
	}))
	t.Log(findSquare([][]int{
		{0},
	}))
	t.Log(findSquare([][]int{
		{1, 1, 1, 0, 1, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 0, 0, 0, 1, 1},
		{0, 0, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
	}))
	t.Log(findSquare([][]int{
		{1, 1, 1, 1, 0, 1, 0, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 1, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 1, 1, 0, 1, 0, 1},
		{0, 0, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 0, 1, 1, 1},
	}))
	/*
		0, 0, 0
		0, 0, 1
		0, 1, 0

		0, 0, 0
		0, 1, 0
		0, 0, 0
	*/
}

func findSquare(matrix [][]int) []int {
	if len(matrix) == 1 {
		if matrix[0][0] == 1 {
			return nil
		}
		return []int{0, 0, 1}
	}
	ret := make([]int, 3)
	x := len(matrix)
	dp := make([][]int, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]int, x)
	}
	dp[x-1][x-1] = arrived(matrix[x-1][x-1])
	for i := x - 2; i >= 0; i-- {
		dp[i][x-1] = arrived(matrix[i][x-1])
	}
	for i := x - 2; i >= 0; i-- {
		dp[x-1][i] = arrived(matrix[x-1][i])
	}
	max := 0
	for i := x - 2; i >= 0; i-- {
		for j := x - 2; j >= 0; j-- {
			a := arrived(matrix[i][j])
			if a == 1 {
				a = min3(dp[i+1][j], dp[i][j+1], dp[i+1][j+1]) + 1
			}
			if max <= a {
				max = a
				ret[0] = i
				ret[1] = j
				ret[2] = a
			}
			dp[i][j] = a
		}
	}
	if max == 0 {
		return nil
	}
	return ret
}

func arrived(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}
