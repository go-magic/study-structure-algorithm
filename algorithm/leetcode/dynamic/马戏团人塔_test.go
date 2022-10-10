package dynamic

import "testing"

func TestBestSeqAtIndex(t *testing.T) {
	if num := bestSeqAtIndex([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110}); num != 6 {
		t.Fatal("请求失败:", num)
		return
	}
}

/*
f(i)=x
*/
func bestSeqAtIndex(height []int, weight []int) int {
	return 0
}


