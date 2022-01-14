package dynamic

import "testing"

func TestNthUglyNumber(t *testing.T) {
	if n := nthUglyNumber(2); n != 2 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(3); n != 3 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(4); n != 4 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(5); n != 5 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(6); n != 6 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(7); n != 8 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(10); n != 12 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(11); n != 15 {
		t.Fatal(n)
		return
	}
	if n := nthUglyNumber(1352); n != 15 {
		t.Fatal(n)
		return
	}
	t.Log("测试成功")
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([][3]int, n)
	dp[0][0] = 0
	saver := make(map[int]struct{})
	saver[1] = struct{}{}
	saver[2] = struct{}{}
	saver[3] = struct{}{}
	saver[5] = struct{}{}
	for i := 1; i < n; i++ {
		d2 := add2(dp[i-1][0], saver)
		d3 := add3(dp[i-1][1], saver)
		d5 := add5(dp[i-1][2], saver)
		dp[i][0] = dp[i-1][0]
		dp[i][1] = dp[i-1][1]
		dp[i][2] = dp[i-1][2]
		minNum := min(min(d2, d3), d5)
		saver[minNum] = struct{}{}
		if d2 == minNum {
			dp[i][0] = d2
		}
		if d3 == minNum {
			dp[i][1] = d3
		}
		if d5 == minNum {
			dp[i][2] = d5
		}
	}
	return max(max(dp[n-1][0], dp[n-1][1]), dp[n-1][2])
}

func add2(value int, saver map[int]struct{}) int {
	temp := value + 2
	for {
		_, exist := saver[temp/2]
		if exist {
			return temp
		}
		temp = temp + 2
	}
}

func add3(value int, saver map[int]struct{}) int {
	temp := value + 3
	for {
		_, exist := saver[temp/3]
		if exist {
			return temp
		}
		temp = temp + 3
	}
}

func add5(value int, saver map[int]struct{}) int {
	temp := value + 5
	for {
		_, exist := saver[temp/5]
		if exist {
			return temp
		}
		temp = temp + 5
	}
}
