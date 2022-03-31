package main

import "algorithm-go/tree"

// 671. 二叉树中第二小的节点

func findSecondMinimumValue(root *tree.TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return -1
	}
	// 左右子节点中不同于root.val的那个值可能是第二小的值
	left, right := root.Left.Val, root.Right.Val
	if root.Val == root.Left.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == root.Right.Val {
		right = findSecondMinimumValue(root.Right)
	}
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	return min(left, right)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
