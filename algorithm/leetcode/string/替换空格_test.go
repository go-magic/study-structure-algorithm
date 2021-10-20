package string

import "testing"

func replaceSpace(s string) string {
	bt := []byte(s)
	str := ""
	for _, v := range bt {
		if v != ' ' {
			str += string(v)
		} else {
			str += `%20`
		}
	}
	return str
}

func TestReplaceString(t *testing.T) {
	str := replaceSpace("We are happy.")
	if str == "We%20are%20happy." {
		t.Log("测试通过")
		return
	}
	t.Fatal(str)
}
