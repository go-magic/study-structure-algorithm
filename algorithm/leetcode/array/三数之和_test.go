package array

import "testing"

func TestCreateArrays(t *testing.T) {

}

func TestThreeSum(t *testing.T) {

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
