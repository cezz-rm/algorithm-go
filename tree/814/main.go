package main

import "algorithm-go/tree"

// 814. 二叉树剪枝

func pruneTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
