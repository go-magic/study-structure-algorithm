package dynamic

import (
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	nodes := generateTrees(3)
	t.Log("测试通过", nodes)
	nodes = generateTrees(1)
	t.Log("测试通过", nodes)
	t.Log("测试通过")
}

/*

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

输入：n = 1
输出：[[1]]
*/
func generateTrees(n int) []*TreeNode {
	return []*TreeNode{}
}
