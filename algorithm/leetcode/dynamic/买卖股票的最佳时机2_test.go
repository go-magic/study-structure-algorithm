package dynamic

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
	if num := maxProfit2([]int{2,1,2,0,1}); num != 2 {
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
	initMaxProfitSaver2(len(prices))
	return buyMaxProfit2(-1,0,prices)
}

func buyMaxProfit2(value,index int ,prices []int) int {
	if index == len(prices) {
		return 0
	}
	if value == -1 {
		s := buyMaxProfit2(prices[index],index+1,prices)
		ns := getMaxProfitSaver2(index+1,prices)
		if s > ns {
			return s
		}
		return ns
	}
	s := getMaxProfitSaver2(index+1,prices) + prices[index] - value
	ns := buyMaxProfit2(value,index+1,prices)
	if s > ns {
		return s
	}
	return ns
}
var maxProfitSaver2 []int

func initMaxProfitSaver2(n int){
	maxProfitSaver2 = make([]int,0,n)
	for i := 0; i < n;i++ {
		maxProfitSaver2 =append(maxProfitSaver2,-1)
	}
}

func getMaxProfitSaver2(index int,prices []int)int{
	if index == len(prices) {
		return 0
	}
	saver :=  maxProfitSaver2[index]
	if saver == -1 {
		saver =  buyMaxProfit2(-1,index,prices)
		maxProfitSaver2[index] = saver
	}
	return saver
}
