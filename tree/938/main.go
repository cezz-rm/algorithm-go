package main

import "algorithm-go/tree"

// 938. 二叉搜索树的范围和

func rangeSumBST(root *tree.TreeNode, low, high int) int {
	var res int
	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		if node.Val < low {
			dfs(node.Right)
		} else if node.Val > high {
			dfs(node.Left)
		} else {
			res += node.Val
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return res
}
