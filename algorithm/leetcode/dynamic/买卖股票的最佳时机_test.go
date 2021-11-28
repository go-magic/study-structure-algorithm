package dynamic

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	if num := maxProfit([]int{7, 1, 5, 3, 6, 4}); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit([]int{7, 6, 4, 3, 1}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

var maxProfitSaver []int

func maxProfit(prices []int) int {
	length := len(prices)
	maxValue := 0
	initMaxArray(prices)
	for i := 0; i < length; i++ {
		temp := maxProfitSaver[i] - prices[i]
		if temp > maxValue {
			maxValue = temp
		}
	}
	return maxValue
}

func initMaxArray(prices []int) {
	length := len(prices)
	maxProfitSaver = make([]int, length)
	maxSaver := 0
	for i := length - 1; i >= 0; i-- {
		temp := prices[i]
		if maxSaver < temp {
			maxSaver = temp
		}
		maxProfitSaver[i] = maxSaver
	}
}
