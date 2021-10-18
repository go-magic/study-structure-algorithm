package array

import "testing"

func threeSum(nums []int) [][]int {
	return getAll(nums)
}

func createArr(length int) [][]int {
	arrs := make([][]int, 0, length-1)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length; j++ {
			arr := make([]int, length)
			arrs = append(arrs, arr)
		}
	}
	return arrs
}

func getAll(nums []int) [][]int {
	arrs := make([][]int, 0, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		arri := createArr(len(nums))
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				arri[j][k] = nums[k]
			}
		}
		arrs = append(arrs, arri...)
	}
	return arrs
}

func fullArr(sum []int, arrs [][]int) [][]int {
	for i := 0; i < len(sum)-1; i++ {
		for j := 0; j < len(sum); j++ {
			for k := j; k < len(sum); k++ {

			}
		}
	}
}

func TestCreateArrays(t *testing.T) {
	arrs := createArr(3)
	t.Log(arrs)
}

func TestThreeSum(t *testing.T) {
	arrs := threeSum([]int{1, 2, 3})
	t.Log(arrs)
}

/*
1,2,3总共6种排列组合(每条分支一种)
	 1					2                3
   2   3			  1   3            1    2
 3	     2          3       1        2        1


 1,2,3,4总共24种排列组合(每条分支一种)
	1
	2     3      4
	3  4  2  4   2  3
	4  3  4  2   3  2

	2
	1     3      4
	3  4  1  4   1  3
	4  3  4  1   3  1

	3
	1     2      4
	2  4  1  4   1  2
	4  2  4  1   2  1

	4
	1     2      3
	2  3  1  3   1  2
	3  2  3  1   2  1


	针对重复优化【1，2，3，3】

	1
	2  3
	3  2

	2
	1  3
	3  1

	3
	1  2
	2  1
*/
