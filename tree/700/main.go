package main

import (
	"algorithm-go/tree"
)

// 700. 二叉搜索树中的搜索

func searchBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}
