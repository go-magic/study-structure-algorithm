package node

import "testing"

func TestBNodeInsert(t *testing.T) {
	b := NewBlock(10)
	if b.Head == nil || b.Head.Data != 10 {
		t.Fatal("10", b)
		return
	}
	b.Insert(9)
	if b.Head == nil || b.Head.Data != 9 || b.Head.Next == nil || b.Head.Next.Data != 10 {
		t.Fatal("10,9", b)
		return
	}
	b.Insert(10)
	if b.Head == nil || b.Head.Data != 9 || b.Head.Next == nil || b.Head.Next.Data != 10 {
		t.Fatal("10,9,10", b)
		return
	}
	b.Insert(11)
	if b.Head == nil || b.Head.Data != 9 ||
		b.Head.Next == nil || b.Head.Next.Data != 10 ||
		b.Head.Next.Next == nil || b.Head.Next.Next.Data != 11 {
		t.Fatal("10,9,10,11", b)
		return
	}
	b.Insert(12)
	if b.Head == nil || b.Head.Data != 9 ||
		b.Head.Next == nil || b.Head.Next.Data != 10 ||
		b.Head.Next.Next == nil || b.Head.Next.Next.Data != 11 ||
		b.Head.Next.Next.Next == nil || b.Head.Next.Next.Next.Data != 12 {
		t.Fatal("10,9,10,11,12", b)
		return
	}
	b.Insert(8)
	if b.Head == nil || b.Head.Data != 8 ||
		b.Head.Next == nil || b.Head.Next.Data != 9 ||
		b.Head.Next.Next == nil || b.Head.Next.Next.Data != 10 ||
		b.Head.Next.Next.Next == nil || b.Head.Next.Next.Next.Data != 11 ||
		b.Head.Next.Next.Next.Next == nil || b.Head.Next.Next.Next.Next.Data != 12 {
		t.Fatal("10,9,10,11,12,8", b)
		return
	}
	p := b.Head
	arr := []int{8, 9, 10, 11, 12}
	for i := 0; i < 5; i++ {
		if p == nil || p.Data != arr[i] {
			t.Fatal("测试失败", p)
			return
		}
		p = p.Next
	}
	t.Log("测试通过")
}
