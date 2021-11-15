package dynamic

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	pmap = make(map[int]Pkg)
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
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
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
	pmap = make(map[int]Pkg)
	s := getMaxSubArrayIndex1(0, 4, []int{4, -1, 2, 1, -5})
	t.Log(s)
}

type Pkg struct {
	Left, Right, Max int
}

var pmap map[int]Pkg

func getMaxSubArrayIndex1(max, index int, arr []int) Pkg {
	if index == 0 {
		return Pkg{Left: arr[0] + max, Right: arr[0]}
	}
	count++
	max = max + arr[index]
	pkg, ok := pmap[index]
	if !ok {
		pkg = getMaxSubArrayIndex1(0, index-1, arr)
		pkg.Right = pkg.Right + arr[index]
		pmap[index] = pkg
	}
	pkg.Right = max
	temp := maxValue(pkg.Left, pkg.Right, max)
	pkg.Right = pkg.Left
	if temp > pkg.Max {
		pkg.Max = temp
	}
	return pkg
}

func maxValue(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
