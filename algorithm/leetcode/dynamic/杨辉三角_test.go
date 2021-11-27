package dynamic

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	arr := generate(5)
	t.Log(arr)
	arr = generate(1)
	t.Log(arr)
	t.Log("测试通过")
}

func generate(numRows int) [][]int {
	rets := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		ret := make([]int, i+1)
		ret[0] = 1
		ret[i] = 1
		for j := 1; j < i; j++ {
			ret[j] = rets[i-1][j-1] + rets[i-1][j]
		}
		rets = append(rets, ret)
	}
	return rets
}

/*
f(x,y) = f(x-1,y-1)+f(x-1,y)
*/
