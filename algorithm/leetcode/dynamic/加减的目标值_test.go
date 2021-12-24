package dynamic

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	if num := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3); num != 5 {
		t.Fatal("测试不通过")
		return
	}
	if num := findTargetSumWays([]int{1}, 1); num != 1 {
		t.Fatal("测试不通过")
		return
	}
	if num := findTargetSumWays([]int{1, 0}, 1); num != 2 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

var findTargetNums []int
var saverFindTargetNums [][]int

func findTargetSumWays(nums []int, target int) int {
	saverFindTargetNums = make([][]int, 40002)
	for i := 0; i < 40002; i++ {
		arr := make([]int, 21)
		for j := 0; j < 21; j++ {
			arr[j] = 20002
		}
		saverFindTargetNums[i] = arr
	}
	findTargetNums = nums
	return findTarget(0, 0, target)
}

func findTarget(now, index, target int) int {
	if index == len(findTargetNums) {
		if now == target {
			return 1
		}
		return 0
	}
	sum := 0
	add := saverFindTargetNums[20001+now+findTargetNums[index]][index+1]
	if add == 20002 {
		add = findTarget(now+findTargetNums[index], index+1, target)
	}
	sub := saverFindTargetNums[20001+now+findTargetNums[index]][index+1]
	if sub == 20002 {
		sub = findTarget(now-findTargetNums[index], index+1, target)
	}
	sum = add + sub
	return sum
}

//leetcode解答
func findTargetSumWays1(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < target || (sum-target)&1 == 1 {
		return 0
	}
	neg := (sum - target) / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}
