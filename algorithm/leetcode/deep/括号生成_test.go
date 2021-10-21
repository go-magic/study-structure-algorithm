package deep

import "testing"

func generateParenthesis(n int) []string {
	return deep(n*2, 0, []string{}, "")
}

func deep(x, y int, ret []string, s string) []string {
	if y == x {
		ret = append(ret, s)
		return ret
	}
	s = insert(s)
	for i := 0; i < x; i++ {
		ret = deep(x, i, ret, s)
	}
	return ret
}

func insert(s string) string {

}

func deep1(n int) {
	for i := 0; i < n*2-1; i++ {

	}
}

/*
case 1:
	input 3
	output ["((()))","(()())","(())()","()(())","()()()"]
case 2:
	input 1
	output ["()"]
*/
func TestGenerateParenthesis(t *testing.T) {
	arr := generateParenthesis(1)
	if arr[0] != "()" {
		t.Fatal("测试失败")
		return
	}
	arr = generateParenthesis(3)
	t.Log("测试通过")
}
