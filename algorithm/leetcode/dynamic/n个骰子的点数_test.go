package dynamic

import "testing"

func TestDicesProbability(t *testing.T) {
	arr := dicesProbability(1)
	t.Log(arr)
	t.Log(len(arr))
}

func dicesProbability(n int) []float64 {
	length := n * 6
	ret := make([]float64, 0, length)
	saver := make([]int, length)
	var base = 0
	for i := 0; i < n; i++ {
		base = base + 6
		for j := i; j < 6*i; j++ {
			saver[j] = saver[j] + 1
			saver[j+1] = saver[j+1] + 2
			saver[j+2] = saver[j+2] + 3
			saver[j+3] = saver[j+3] + 4
			saver[j+4] = saver[j+4] + 5
			saver[j+5] = saver[j+5] + 6
		}
	}
	per := float64(1) / 11
	for i := 0; i < n*6-n+1; i++ {
		ret = append(ret, per*float64(saver[i]))
	}
	return ret
}
