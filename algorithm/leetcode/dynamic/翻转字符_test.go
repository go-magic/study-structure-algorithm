package dynamic

import (
	"testing"
)

func TestMinFlipsMonoIncr(t *testing.T) {
	if num := minFlipsMonoIncr("00110"); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := minFlipsMonoIncr("010110"); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	if num := minFlipsMonoIncr("00011000"); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
f(n)=f()
*/
func minFlipsMonoIncr(s string) int {
	length := len(s)
	dp := make([][2]int, length)
	if s[0] == '0' {
		dp[0][0] = 0
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
		dp[0][1] = 0
	}
	for i := 1; i < length; i++ {
		if s[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[len(s)-1][0], dp[len(s)-1][1])
}
