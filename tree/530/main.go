package main

import (
	"algorithm-go/tree"
	"math"
)

// 530. 二叉搜索树的最小绝对差

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinimumDifference(root *tree.TreeNode) int {
	var dfs func(node *tree.TreeNode)
	var prev *tree.TreeNode
	minNum := math.MaxInt32
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			minNum = min(minNum, int(math.Abs(float64(node.Val-prev.Val))))
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return minNum
}
