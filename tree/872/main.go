package main

import "algorithm-go/tree"

// 872. 叶子相似的树

func leafSimilar(root1 *tree.TreeNode, root2 *tree.TreeNode) bool {
	var vals []int
	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			vals = append(vals, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	vals1 := append([]int{}, vals...)
	vals = []int{}
	dfs(root2)
	if len(vals) != len(vals1) {
		return false
	}
	for i, v := range vals1 {
		if v != vals[i] {
			return false
		}
	}
	return true
}
