package greed

import (
	"testing"
)

func TestMaxProfit2(t *testing.T) {
	if num := maxProfit2([]int{7,6,4,3,1}); num != 0 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit2([]int{1,2,3,4,5}); num != 4 {
		t.Fatal("测试不通过")
		return
	}
	if num := maxProfit2([]int{7,1,5,3,6,4}); num != 7 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

*/
func maxProfit2(prices []int) int {
	temp := prices[0]
	value := 0
	length := len(prices)
	for i := 1; i < length;i++ {
		if temp > prices[i] {
			temp = prices[i]
		}else {
			value += prices[i] - temp
			temp = prices[i]
		}
	}
	return value
}
