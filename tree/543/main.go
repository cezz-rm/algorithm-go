package main

import "algorithm-go/tree"

// 543. 二叉树的直径
// 一棵二叉树的直径长度是任意两个节点长度的最大值

func diameterOfBinaryTree(root *tree.TreeNode) int {
	var maxDiameter int
	var maxDepth func(node *tree.TreeNode) int
	maxDepth = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := maxDepth(node.Left)
		rightMax := maxDepth(node.Right)
		maxDiameter = max(maxDiameter, leftMax+rightMax)
		return 1 + max(leftMax, rightMax)
	}
	maxDepth(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
