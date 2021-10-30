package dynamic

import (
	"fmt"
	"testing"
)

func TestCoinChange1(t *testing.T) {
	//if num := coinChange1([]int{1, 2, 5}, 11); num != 3 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange1([]int{2}, 3); num != -1 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange1([]int{1}, 0); num != 0 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange1([]int{1}, 1); num != 1 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange1([]int{1}, 2); num != 2 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	//count = 0
	//if num := coinChange1([]int{416, 157, 454, 343}, 1761); num != -1 {
	//	t.Fatal("测试不通过")
	//	return
	//}
	//t.Log(count)
	count = 0
	if num := coinChange1([]int{186, 419, 83, 408}, 6249); num != 20 {
		t.Fatal("测试不通过")
		return
	}
	t.Log(count)
	t.Log("测试通过")
}

func coinChange1(coins []int, amount int) int {
	counts = initArray(amount)
	num, ok := minCoin1(coins, amount)
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

var counts []int

var validMap map[int]bool

func minCoin1(coins []int, amount int) (int, bool) {
	if amount == 0 {
		return 0, true
	}
	if !valid(coins, amount) {
		return 0, false
	}
	min := 2 << 32
	var isok bool
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		temp := counts[amount-coin]
		ok := false
		if temp == 0 {
			count++
			temp, ok = minCoin1(coins, amount-coin)
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
	fmt.Printf("写入%d,%d,%d\n", amount, min, count)
	counts[amount] = min
	//失败的也应该存下来
	//if isok {
	//	m[amount] = min
	//	fmt.Printf("写入%d,值为:%d,count:%d\n", amount, min, count)
	//}
	return min, isok
}

func initArray(amount int) []int {
	validMap = make(map[int]bool)
	arr := make([]int, amount+1)
	return arr
}
