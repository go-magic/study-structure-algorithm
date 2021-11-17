package dynamic

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	if value := lengthOfLongestSubstring(" "); value != 1 {
		t.Fatal(value)
		return
	}
	/*
		1, 3, 1
	*/
	if value := lengthOfLongestSubstring("abcabcb"); value != 3 {
		t.Fatal(value)
		return
	}
	/*
		1, 3, 1
		1, 5, 1
	*/
	if value := lengthOfLongestSubstring("bbbbb"); value != 1 {
		t.Fatal(value)
		return
	}
	/*
		1, 3, 1
		1, 5, 1
		4, 2, 1
	*/
	if value := lengthOfLongestSubstring("pwwkew"); value != 3 {
		t.Fatal(value)
		return
	}
	if value := lengthOfLongestSubstring("abcaadceef"); value != 4 {
		t.Fatal(value)
		return
	}
	t.Log("测试通过")
}

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
*/
func lengthOfLongestSubstring(s string) int {
	arr = initArr(256)
	max = 0
	return getMaxLengthSubstring(s)
}

/*
pwwkew
f(0) = max(1+f(1),f(1))
*/
func getMaxLengthSubstring(s string) int {
	length := len(s)
	temp := 0
	index := 0
	for i := 0; i < length; i++ {
		at := int(s[i])
		if a := arr[at]; a != -1 {
			index = i - arr[at]
			if index <= temp {
				temp = index
			} else {
				temp++
				if temp > max {
					max = temp
				}
			}
		} else {
			temp++
			if temp > max {
				max = temp
			}
		}
		arr[at] = i
	}
	return max
}

func initArr(num int) []int {
	ret := make([]int, num)
	for i := 0; i < 256; i++ {
		ret[i] = -1
	}
	return ret
}

var arr = []int{}
var max int
