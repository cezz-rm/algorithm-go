package main

import "algorithm-go/tree"

// 101. 对称二叉树

func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *tree.TreeNode) bool {
	if left == nil && right == nil {
		return left == right
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
