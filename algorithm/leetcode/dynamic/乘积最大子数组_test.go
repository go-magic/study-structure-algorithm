package dynamic

import "testing"

func TestMaxProduct(t *testing.T) {
	if num := maxProduct([]int{2,3,-2,4}); num != 6{
		t.Fatal("测试不通过")
		return
	}
	if num := maxProduct([]int{-2,0,-1}); num !=0 {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
0表示不连续
1表示连续
*/
func maxProduct(nums []int) int {
	length := len(nums)
	rets := make([][3]int,0)
	for i := 1; i <  len(nums);i++{
		rets[i][0] = max3(rets[i-1][0],nums[i])
		rets[i][1] = max3(rets[i-1][1],nums[i])
	}
	return rets[length-1][0]
}
