package main

import (
	"algorithm-go/tree"
	"math"
)

// 783. 二叉搜索树节点最小距离

func minDiffInBST(root *tree.TreeNode) int {
	var res = math.MaxInt32
	var dfs func(node *tree.TreeNode)
	var prev *tree.TreeNode
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			res = min(res, abs(node.Val-prev.Val))
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
