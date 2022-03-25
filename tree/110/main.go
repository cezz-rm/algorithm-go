package main

import (
	"algorithm-go/tree"
	"math"
)

// 110. 平衡二叉树
// 给定一个二叉树, 判断它是否是高度平衡的二叉树

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *tree.TreeNode) bool {
	var isBalance = true
	var maxDepth func(node *tree.TreeNode) int
	maxDepth = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		leftMaxDepth := maxDepth(node.Left)
		rightMaxDepth := maxDepth(node.Right)
		if math.Abs(float64(rightMaxDepth-leftMaxDepth)) > 1 {
			isBalance = false // 后续遍历, 如果左右深度大于1, 就不是平衡二叉树
		}
		return 1 + max(leftMaxDepth, rightMaxDepth)
	}
	maxDepth(root)
	return isBalance
}
