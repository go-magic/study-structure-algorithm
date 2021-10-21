package string

import "testing"

func TestFirstUniqChar(t *testing.T) {
	str := "abaccdeff"
	if a := firstUniqChar(str); a != 'b' {
		t.Fatal("测试失败")
		return
	}
	str = ""
	if a := firstUniqChar(str); a != ' ' {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}

/*
具体看字符有多少个
*/
func firstUniqChar(s string) byte {
	arr := make([]int, 128)
	for _, v := range s {
		arr[v]++
	}
	for _, v := range s {
		if arr[v] == 1 {
			return byte(v)
		}
	}
	return ' '
}
