package queue

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Append(1)
	stack.Append(2)
	stack.Append(3)
	stack.Append(4)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

}

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	queue.AppendTail(4)
	a1 := queue.DeleteHead()
	if a1 != 1 {
		t.Fatal("测试失败1")
		return
	}
	a1 = queue.DeleteHead()
	if a1 != 2 {
		t.Fatal("测试失败2")
		return
	}
	queue.AppendTail(5)
	queue.AppendTail(6)
	a1 = queue.DeleteHead()
	if a1 != 3 {
		t.Fatal(a1)
		return
	}
	a1 = queue.DeleteHead()
	if a1 != 4 {
		t.Fatal("测试失败4")
		return
	}
	a1 = queue.DeleteHead()
	if a1 != 5 {
		t.Fatal("测试失败5")
		return
	}
	a1 = queue.DeleteHead()
	if a1 != 6 {
		t.Fatal("测试失败6")
		return
	}
	t.Log("测试通过")
}
