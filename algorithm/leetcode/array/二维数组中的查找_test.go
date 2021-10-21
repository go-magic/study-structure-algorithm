package array

import "testing"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	valid, _ := find(0, 0, matrix, target-1, target)
	return valid
}

func find(x, y int, matrix [][]int, now, target int) (bool, bool) {
	if now == target {
		return true, false
	}
	if x == len(matrix) {
		return false, false
	}
	if y == len(matrix[x]) {
		return false, false
	}
	if now > target {
		return false, true
	}
	length := len(matrix[x])
	for i := 0; i < length; i++ {
		valid, ok := find(x+1, i, matrix, matrix[x][i], target)
		if valid {
			return true, false
		}
		if ok {
			return false, false
		}
	}
	return false, false
}

func TestFindNumberIn2DArray(t *testing.T) {
	arrs := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if ok := findNumberIn2DArray(arrs, 5); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := findNumberIn2DArray(arrs, 20); ok {
		t.Fatal("测试不通过")
		return
	}
	arrs = [][]int{{-1, 3}}
	if ok := findNumberIn2DArray(arrs, 3); !ok {
		t.Fatal("测试不通过")
		return
	}
	arrs = [][]int{{48, 65, 70, 113, 133, 163, 170, 216, 298, 389},
		{89, 169, 215, 222, 250, 348, 379, 426, 469, 554},
		{178, 202, 253, 294, 367, 392, 428, 434, 499, 651},
		{257, 276, 284, 332, 380, 470, 516, 561, 657, 698},
		{275, 331, 391, 432, 500, 595, 602, 673, 758, 783},
		{357, 365, 412, 450, 556, 642, 690, 752, 801, 887},
		{359, 451, 534, 609, 654, 662, 693, 766, 803, 964},
		{390, 484, 614, 669, 684, 711, 767, 804, 857, 1055},
		{400, 515, 683, 732, 812, 834, 880, 930, 1012, 1130},
		{480, 538, 695, 751, 864, 939, 966, 1027, 1089, 1224}}
	if ok := find2(arrs, 929); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*

[
	[48,65,70,113,133,163,170,216,298,389],
	[89,169,215,222,250,348,379,426,469,554],
	[178,202,253,294,367,392,428,434,499,651],
	[257,276,284,332,380,470,516,561,657,698],
	[275,331,391,432,500,595,602,673,758,783],
	[357,365,412,450,556,642,690,752,801,887],
	[359,451,534,609,654,662,693,766,803,964],
	[390,484,614,669,684,711,767,804,857,1055],
	[400,515,683,732,812,834,880,930,1012,1130],
	[480,538,695,751,864,939,966,1027,1089,1224]]
929
*/

/*
暴力求解 n*n
*/
func find1(matrix [][]int, target int) bool {
	for _, val := range matrix {
		for _, v := range val {
			if v == target {
				return true
			}
			if v > target {
				break
			}
		}
	}
	return false
}

/*
线性查找 n+n
*/
func find2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	xMaxLen := len(matrix)
	yMaxLen := len(matrix[0])
	x := 0
	y := yMaxLen - 1
	for x < xMaxLen && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
			continue
		}
		x++
	}
	return false
}
