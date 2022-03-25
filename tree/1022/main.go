package main

import (
	"algorithm-go/tree"
)

// 1022. 从根到叶的二进制数之和

func sumRootToLeaf(root *tree.TreeNode) int {
	var res int // 总和
	var sum int // 单条路径和

	var dfs func(node *tree.TreeNode, sum int)
	dfs = func(node *tree.TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			res += sum
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, sum)
	return res
}
