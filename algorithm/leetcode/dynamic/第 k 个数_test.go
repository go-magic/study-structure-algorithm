package dynamic

import (
	"testing"
)

func TestGetKthMagicNumber(t *testing.T) {
	if num := getKthMagicNumber(5); num != 9 {
		t.Fatal("测试失败:", num)
		return
	}
	if num := getKthMagicNumber(6); num != 15 {
		t.Fatal("测试失败:", num)
		return
	}
	if num := getKthMagicNumber(7); num != 21 {
		t.Fatal("测试失败:", num)
		return
	}
	t.Log("测试通过")
}

func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	three, five, seven := 0, 0, 0
	for i := 1; i < k; i++ {
		threeValue := dp[three] * 3
		fiveValue := dp[five] * 5
		sevenValue := dp[seven] * 7
		minValue := min(min(threeValue, fiveValue), sevenValue)
		if minValue == threeValue {
			three++
		}
		if minValue == fiveValue {
			five++
		}
		if minValue == sevenValue {
			seven++
		}
		dp[i] = minValue
	}
	return dp[k-1]
}
