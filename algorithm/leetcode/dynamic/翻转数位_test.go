package dynamic

import (
	"testing"
)

/*
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

示例 1：

输入: num = 1775(11011101111)
输出: 8
示例 2：

输入: num = 7(111)
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-bits-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseBits(num int) int {
	max := ReverseBits(num)
	if max > 32 {
		max = 32
	}
	return max
}

func ReverseBits(num int) int {
	if num == 0 {
		return 1
	}
	bits := IntToBits(num)
	zeros := make([]int, 0)
	for i := 0; i < len(bits); i++ {
		if bits[i] == 0 {
			zeros = append(zeros, i-1)
		}
		if i == len(bits)-1 && bits[i] == 1 {
			zeros = append(zeros, i)
		}
	}
	if len(zeros) == 1 {
		if bits[len(bits)-1] == 0 {
			return len(bits)
		}
		return len(bits) + 1
	}
	if len(zeros) == 2 {
		if bits[len(bits)-1] == 0 {
			return len(bits) - 1
		}
		return len(bits)
	}
	max := zeros[1] + 1
	for i := 2; i < len(zeros); i++ {
		tmp := zeros[i] - zeros[i-2] - 1
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func IntToBits(num int) []int {
	if num < 0 {
		num = 0xFFFFFFFF + num + 1
	}
	bits := make([]int, 0)
	for num != 0 {
		bits = append(bits, num%2)
		num = num / 2
	}
	ret := make([]int, 0, len(bits))
	for i := len(bits) - 1; i >= 0; i-- {
		ret = append(ret, bits[i])
	}
	return ret
}

func Test123(t *testing.T) {
	t.Log(0xFFFFFFFF | 7)
}

func TestReverseBits(t *testing.T) {
	t.Log(IntToBits(1775))
	if length := reverseBits(1775); length != 8 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(7); length != 4 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(11); length != 4 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(14); length != 4 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(58); length != 5 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(19); length != 3 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(0); length != 1 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(1); length != 2 {
		t.Fatal("测试失败")
		return
	}
	if length := reverseBits(-1); length != 32 {
		t.Fatal("测试失败")
		return
	}
	t.Log("测试通过")
}
