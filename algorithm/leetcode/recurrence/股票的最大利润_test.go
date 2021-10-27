package recurrence

import "testing"

func TestMaxProfit(t *testing.T) {
	if max := maxProfit([]int{7, 1, 5, 3, 6, 4}); max != 5 {
		t.Fatal("测试失败")
		return
	}
	if max := maxProfit([]int{7, 6, 4, 3, 1}); max != 0 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

var maxTemp int

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	min := prices[0]
	max := 0
	for i := 0; i < length; i++ {
		if min < prices[i] {
			temp := prices[i] - min
			if temp > max {
				max = temp
			}
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return max
}

func find(val, pos int, prices []int) {
	if pos == len(prices) {
		return
	}
	if val < prices[pos] {
		length := len(prices)
		tempMax := prices[pos] - val
		for i := pos; i < length; i++ {
			m := prices[i] - val
			if m > tempMax {
				tempMax = m
			}
		}
		if maxTemp < tempMax {
			maxTemp = tempMax
		}
		val = prices[pos]
	}
	if val > prices[pos] {
		val = prices[pos]
	}
	find(val, pos+1, prices)
}
