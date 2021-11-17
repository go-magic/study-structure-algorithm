package dynamic

import (
	"testing"
)

func TestTranslateNum(t *testing.T) {
	if num := translateNum(1); num != 1 {
		t.Fatal(num)
		return
	}
	if num := translateNum(12); num != 2 {
		t.Fatal(num)
		return
	}
	if num := translateNum(123); num != 3 {
		t.Fatal(num)
		return
	}
	if num := translateNum(12258); num != 5 {
		t.Fatal(num)
		return
	}
	if num := translateNum(624); num != 2 {
		t.Fatal(num)
		return
	}
	if num := translateNum(506); num != 1 {
		t.Fatal(num)
		return
	}
	if num := translateNum(18580); num != 2 {
		t.Fatal(num)
		return
	}
	t.Log("测试通过")
}

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/
func translateNum(num int) int {
	if num/10 == 0 {
		return 1
	}
	translateMapOne = make(map[int]int)
	translateMapTwo = make(map[int]int)
	nums := slice(num)
	one := translate(0, 1, nums)
	two := translate(0, 2, nums)
	return one + two
}

func translate(index, length int, nums []int) int {
	if index+length == len(nums)-1 {
		if !checkTranslateValid(index, length, nums) {
			return 0
		}
		return 1
	}
	if index+length == len(nums) {
		if !checkTranslateValid(index, length, nums) {
			return 0
		}
		return 1
	}
	if !checkTranslateValid(index, length, nums) {
		return 0
	}
	one, exits := translateMapOne[index+length]
	if !exits {
		one = translate(index+length, 1, nums)
		translateMapOne[index+length] = one
	}
	two, exits := translateMapTwo[index+length]
	if !exits {
		two = translate(index+length, 2, nums)
		translateMapTwo[index+length] = two
	}
	return one + two
}

var translateMapOne = map[int]int{}
var translateMapTwo = map[int]int{}

/*
注意*0*这种情况
*/
func checkTranslateValid(index, length int, nums []int) bool {
	num := 0
	for i := 0; i < length; i++ {
		num = num*10 + nums[index+i]
		if num == 0 && length-i == 2 {
			return false
		}
	}
	return num <= 25
}

func slice(num int) []int {
	temp := make([]int, 0)
	for num != 0 {
		temp = append(temp, num%10)
		num = num / 10
	}
	ret := make([]int, 0, len(temp))
	for i := len(temp) - 1; i >= 0; i-- {
		ret = append(ret, temp[i])
	}
	return ret
}
