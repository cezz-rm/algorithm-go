package main

import "algorithm-go/tree"

// 104. 二叉树的最大深度
// 给定一个二叉树, 找出其最大深度

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return 1 + max(leftMax, rightMax)
}

func maxDepth2(root *tree.TreeNode) int {
	depth, res := 0, 0
	var traverse func(node *tree.TreeNode)
	traverse = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		depth++
		res = max(res, depth)
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}
	traverse(root)
	return depth
}
