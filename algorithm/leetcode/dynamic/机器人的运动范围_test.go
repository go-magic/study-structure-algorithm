package dynamic

import "testing"

func TestMovingCount(t *testing.T) {
	if num := movingCount(2, 2, 1); num != 3 {
		t.Fatal("测试不通过")
		return
	}
	if num := movingCount(3, 1, 0); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

func movingCount(m int, n int, k int) int {
	dp = make([][]int, 0, m)
	for i := 0; i < m; i++ {
		d := make([]int, 0, n)
		for j := 0; j < n; j++ {
			d = append(d, -1)
		}
		dp = append(dp, d)
	}
	subMovingCount(0, 0, k)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func countBit(a int) int {
	ret := 0
	for a != 0 {
		temp := a % 10
		ret = ret + temp
		a = a / 10
	}
	return ret
}

func subMovingCount(x, y, k int) {
	if x == len(dp) || y == len(dp[0]) {
		return
	}
	count := countBit(x) + countBit(y)
	if count > k {
		return
	}
	dp[x][y] = 1
	if x == len(dp)-1 && y == len(dp[0])-1 {
		return
	}
	if x < len(dp)-1 && dp[x+1][y] == -1 {
		subMovingCount(x+1, y, k)
	}
	if y < len(dp[0])-1 && dp[x][y+1] == -1 {
		subMovingCount(x, y+1, k)
	}
}

func TestSubMovingCount(t *testing.T) {
	t.Log(countBit(35))
	t.Log(countBit(37))
	if count := movingCount(2, 2, 1); count != 3 {
		t.Fatal("测试失败")
		return
	}
	if count := movingCount(3, 1, 0); count != 1 {
		t.Fatal("测试失败")
		return
	}
	if count := movingCount(2, 3, 17); count != 6 {
		t.Fatal("测试失败")
		return
	}
	if count := movingCount(38, 15, 9); count != 135 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
