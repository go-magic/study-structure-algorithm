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
	//注释掉的最小值都为0,所以可以删除掉
	//if s[0] == '0' {
	//	dp[0][0] = 0
	//	dp[0][1] = 1
	//} else {
	//	dp[0][0] = 1
	//	dp[0][1] = 0
	//}
	for i := 1; i < length; i++ {
		if s[i] == '0' {
			//若需要将当前设置为0,则前面必须全部为0,所以只能从dp[i-1][0]获取值,当前本身为0所以不需要转换
			dp[i][0] = dp[i-1][0]
			//若需要将当前设置为1,则前面可以为0,也可以为1,从两者之间取最小值,当前本身为0所以需要1次转换
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			//若需要将当前设置为0,则前面必须全部为0,所以只能从dp[i-1][0]获取值,当前本身为1所以需要1次转换
			dp[i][0] = dp[i-1][0] + 1
			//若需要将当前设置为1,则前面可以为0,也可以为1,从两者之间取最小值,当前本身为1所以不需要转换
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}
	return min(dp[len(s)-1][0], dp[len(s)-1][1])
}
