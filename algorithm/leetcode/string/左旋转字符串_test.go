package string

import "testing"

func reverseLeftWords(s string, n int) string {
	bt := []byte(s)
	bt = move(bt, n)
	return string(bt)
}

func move(bt []byte, n int) []byte {
	length := len(bt)
	ret := make([]byte, len(bt))
	n = n % length
	if n < 0 {
		n += length
	}
	k := 0
	for i := n; i < length; i++ {
		ret[k] = bt[i]
		k++
	}
	for i := 0; i < n; i++ {
		ret[k] = bt[i]
		k++
	}
	return ret
}

func TestReverseLeftWords(t *testing.T) {
	str1 := reverseLeftWords("abcdefg", 2)
	if str1 != "cdefgab" {
		t.Fatal(str1)
		return
	}
	str1 = reverseLeftWords("lrloseumgh", 6)
	if str1 != "umghlrlose" {
		t.Fatal(str1)
		return
	}
	str1 = reverseLeftWords("abcdefg", -1)
	if str1 != "gabcdef" {
		t.Fatal(str1)
		return
	}
	str1 = reverseLeftWords("abcdefg", 9)
	if str1 != "cdefgab" {
		t.Fatal(str1)
		return
	}
	t.Log("测试通过")
}
