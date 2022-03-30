package main

import (
	"algorithm-go/tree"
	"math"
)

// 563. 二叉树的坡度

func findTilt(root *tree.TreeNode) int {
	var res int
	var dfs func(node *tree.TreeNode) int
	// 定义一棵二叉树, 返回这棵二叉树所有元素的和
	dfs = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res += int(math.Abs(float64(left - right)))
		return left + right + node.Val
	}
	dfs(root)
	return res
}
