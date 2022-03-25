package main

import (
	"algorithm-go/tree"
)

// 404. 左叶子之和
// 给定二叉树的根节点root, 返回所有左叶子之和

func sumOfLeftLeaves(root *tree.TreeNode) int {
	var sum int
	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sum
}
