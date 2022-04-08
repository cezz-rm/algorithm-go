package main

import (
	"algorithm-go/tree"
)

// 965. 单值二叉树

func isUnivalTree(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
