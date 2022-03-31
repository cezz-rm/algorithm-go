package main

import "algorithm-go/tree"

// 653.两数之和IV

func findTarget(root *tree.TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(node *tree.TreeNode) bool
	dfs = func(node *tree.TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
