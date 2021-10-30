package dynamic

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	//if num := coinChange([]int{1, 2, 5}, 11); num != 3 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange([]int{2}, 3); num != -1 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange([]int{1}, 0); num != 0 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange([]int{1}, 1); num != 1 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange([]int{1}, 2); num != 2 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	count = 0
	if num := coinChange([]int{416, 157, 454, 343}, 1761); num != -1 {
		t.Fatal("测试不通过")
		return
	}
	t.Log(count)
	count = 0
	if num := coinChange([]int{186, 419, 83, 408}, 6249); num != 20 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
*/
func coinChange(coins []int, amount int) int {
	m = make(map[int]int)
	validMap = make(map[int]bool)
	num, ok := minCoin(coins, amount)
	if !ok {
		return -1
	}
	return num
}

/*
	[2,5,7]
	f(27) = x
	f(27-2) = x-1
	f(27-5) = x-1
	f(27-7) = x-1

	min()
*/

var m map[int]int

var count int

func minCoin(coins []int, amount int) (int, bool) {
	if amount == 0 {
		return 0, true
	}
	if !valid(coins, amount) {
		return 0, false
	}
	min := 2 << 32
	var isok bool
	for _, coin := range coins {
		temp, ok := m[amount-coin]
		if !ok {
			count++
			temp, ok = minCoin(coins, amount-coin)
			if !ok {
				continue
			}
		}
		if temp == -1 {
			continue
		}
		temp = temp + 1
		if temp < min {
			min = temp
		}
		isok = true
	}
	if !isok {
		min = -1
	}
	m[amount] = min
	fmt.Printf("写入%d,%d,%d\n", amount, min, count)
	//失败的也应该存下来
	//if isok {
	//	m[amount] = min
	//	fmt.Printf("写入%d,值为:%d,count:%d\n", amount, min, count)
	//}
	return min, isok
}

func valid(coins []int, amount int) bool {
	ok, exist := validMap[amount]
	if exist {
		return ok
	}
	for _, v := range coins {
		if amount >= v {
			validMap[amount] = true
			return true
		}
	}
	validMap[amount] = false
	return false
}
