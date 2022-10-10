package array

import "testing"

var baseMagicNumbers = []int{3,5,7}

func getNextMagicNumber()  {
	
}

func getKthMagicNumber(k int) int {
	return -1
}

func TestGetKthMagicNumber(t *testing.T) {
	cases := map[int]int{1:1,2:3,3:5,4:7,5:9,6:15,7:21}
	for index,hope := range cases {
		result := getKthMagicNumber(index-1)
		if result != hope {
			t.Fatalf("测试失败,第%d个数,期望值为%d,实际值为%d",index,hope,result)
			return
		}
	}
	t.Log("测试通过")
}
