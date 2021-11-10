package dynamic

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	if num := maxSubArray([]int{-2, 1}); num != 1 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3}); num != 1 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3, 4}); num != 4 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3, 4, -1}); num != 4 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3, 4, -1, 2}); num != 5 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1}); num != 6 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{4, -1, 2, 1, -5}); num != 6 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{4, -1, -1, 2, 1, -5}); num != 5 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{4, -1, 1, 2, 1, -5}); num != 7 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	if num := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); num != 6 {
		t.Fatal(num)
		return
	}
	t.Log(count)
	count = 0
	t.Log("测试通过")
}

var smap map[int]int
var nsmap map[int]int

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	smap = make(map[int]int)
	nsmap = make(map[int]int)
	return getMaxSubArrayIndex(0, len(nums)-1, nums)
}

/*
f(n) = max(f(n-1) + arr[n],f(n-1))
*/
func getMaxSubArray(max, index int, arr []int) int {
	if index == 0 {
		return max + arr[0]
	}
	var (
		s  int
		ns int
		ok bool
	)
	max = max + arr[index]
	if s, ok = smap[index]; !ok {
		s = getMaxSubArray(max, index-1, arr)
		smap[index] = s
	}
	if ns, ok = nsmap[index]; !ok {
		ns = getMaxSubArray(0, index-1, arr)
		nsmap[index] = ns
	}
	if s > ns {
		if s > max {
			return s
		}
		return max
	}
	if ns > max {
		return ns
	}
	return max
}

func getMaxSubArrayIndex(max, index int, arr []int) int {
	if index == 0 {
		return max + arr[0]
	}
	count++
	max = max + arr[index]
	left := getMaxSubArrayIndex(max, index-1, arr)
	right, ok := nsmap[index-1]
	if !ok {
		count++
		right = getMaxSubArrayIndex(0, index-1, arr)
		nsmap[index-1] = right
	}
	if left > right {
		if left > max {
			return left
		}
		return max
	}
	if right > max {
		return right
	}
	return max
}

func TestGetMaxSubArrayIndex(t *testing.T) {
	s := getMaxSubArrayIndex(0, 4, []int{4, -1, 2, 1, -5})
	t.Log(s)
}
