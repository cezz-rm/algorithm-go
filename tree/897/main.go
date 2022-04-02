package main

import (
	"algorithm-go/tree"
)

// 897.递增顺序搜索树

func increasingBST(root *tree.TreeNode) *tree.TreeNode {
	newRoot := &tree.TreeNode{}
	resNode := newRoot
	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		resNode.Right = node
		node.Left = nil
		resNode = node
		dfs(node.Right)
	}
	dfs(root)
	return newRoot.Right
}
