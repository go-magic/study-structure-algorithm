package dynamic

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {
	if ok := isInterleave("bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa",
		"babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab",
		"babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"); ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("a", "b", "ab"); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("b", "a", "ab"); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("aabcc", "dbbca", "aadbbcbcac"); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("aabcc", "dbbca", "aadbbbaccc"); ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("", "a", "a"); !ok {
		t.Fatal("测试不通过")
		return
	}
	if ok := isInterleave("", "", ""); ok {
		t.Fatal("测试不通过")
		return
	}
	t.Log("测试通过")
}

/*
"bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa"
"babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab"
"babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"
*/

var (
	ss1 string
	ss2 string
	ss3 string
	isInterleaveArr [][]int
)

/*
1.解题思路
*/
func initIsInterleaveArr(x,y int)  {
	isInterleaveArr = make([][]int,0,x)
	for i := 0; i < x+1;i++ {
		arr := make([]int,0,y+1)
		for j := 0; j < y+1;j++ {
			arr = append(arr,-1)
		}
			isInterleaveArr = append(isInterleaveArr,arr)
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	if s1 == "" && s2 == "" {
		return true
	}
	ss1 = s1
	ss2 = s2
	ss3 = s3
	initIsInterleaveArr(len(s1),len(s2))
	var (
		one bool
		two bool
	)
	if s1 != "" {
		one = interleave(1,0,s1[0])
	}
	if s2 != "" {
		two = interleave(0,1,s2[0])
	}
	if one == true || two == true {
		return true
	}
	return false
}

func interleave(x,y int,b byte) bool {
	if x+y == len(ss3) {
		if b != ss3[x+y-1] {
			return false
		}
		return true
	}
	if x == len(ss1) + 1{
		return false
	}
	if y == len(ss2) +1 {
		return false
	}
	if b != ss3[x+y-1] {
		return false
	}
	var (
		one,two bool
	)
	if x < len(ss1){
		flag := isInterleaveArr[x+1][y]
		if flag == -1 {
			one = interleave(x+1,y,ss1[x])
			isInterleaveArr[x+1][y] = 0
			if one == true {
				isInterleaveArr[x+1][y] = 1
			}
		}else {
			one = translateFlag(flag)
		}
	}
	if y < len(ss2){
		flag := isInterleaveArr[x][y+1]
		if flag == -1 {
			two = interleave(x,y+1,ss2[y])
			isInterleaveArr[x][y+1] = 0
			if two == true {
				isInterleaveArr[x][y+1] = 1
			}
		}else {
			two = translateFlag(flag)
		}
	}
	if one == true || two == true {
		return true
	}
	return false
}

func translateFlag(flag int) bool {
	if flag == 1 {
		return true
	}
	return false
}
