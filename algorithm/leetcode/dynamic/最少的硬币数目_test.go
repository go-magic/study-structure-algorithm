package dynamic

import (
	"testing"
)

func TestCoinChange2(t *testing.T) {
	if num := coinChange2([]int{1, 2, 5}, 11); num != 3 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := coinChange2([]int{2}, 3); num != -1 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := coinChange2([]int{1}, 0); num != 0 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := coinChange2([]int{1}, 1); num != 1 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := coinChange2([]int{1}, 2); num != 2 {
		t.Fatal("测试不通过:", num)
		return
	}
	if num := coinChange2([]int{186, 419, 83, 408}, 6249); num != 2 {
		t.Fatal("测试不通过:", num)
		return
	}
	t.Log("测试通过")
}

var coinSave []int

func coinChange2(coins []int, amount int) int {
	coinSave = make([]int, amount)
	for i := 0; i < amount; i++ {
		coinSave[i] = -2
	}
	return coin(coins, amount)
}

func coin(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	min := 10001
	for _, v := range coins {
		if amount-v < 0 {
			continue
		}
		temp := coinSave[amount-v]
		if temp == -2 {
			temp = coin(coins, amount-v)
			coinSave[amount-v] = temp
		}
		if temp == -1 {
			continue
		}
		temp++
		if temp < min {
			min = temp
		}
	}
	if min == 10001 {
		min = -1
	}
	return min
}
