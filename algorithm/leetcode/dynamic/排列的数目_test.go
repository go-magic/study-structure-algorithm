package dynamic

import (
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	//if num := combinationSum4([]int{1, 2, 3}, 4); num != 7 {
	//	t.Fatal("测试不通过:", num)
	//	return
	//}
	//if num := combinationSum4([]int{2}, 3); num != 0 {
	//	t.Fatal("测试不通过:", num)
	//	return
	//}
	//if num := combinationSum4([]int{2, 1}, 3); num != 3 {
	//	t.Fatal("测试不通过:", num)
	//	return
	//}
	//if num := combinationSum4([]int{1, 50}, 200); num != 28730 {
	//	t.Fatal("测试不通过:", num)
	//	return
	//}
	if num := combinationSum4([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130,
		140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300,
		310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470,
		480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640,
		650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810,
		820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990, 111}, 999); num != 1 {
		t.Fatal("测试不通过:", num)
		return
	}
	t.Log("测试通过")
}

var targets []int
var dpCombination = [2000]int{}

func combinationSum4(nums []int, target int) int {
	targets = nums
	for i := 0; i < 2000; i++ {
		dpCombination[i] = -1
	}
	return combination(0, target)
}

func combination1(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func combination(now, target int) int {
	if now > target {
		return 0
	}
	if now == target {
		return 1
	}
	sum := 0
	for _, v := range targets {
		temp := dpCombination[now+v]
		if temp == -1 {
			temp = combination(now+v, target)
			dpCombination[now+v] = temp
		}
		sum += temp
	}
	return sum
}
