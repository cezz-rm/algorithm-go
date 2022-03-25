package main

import "algorithm-go/tree"

// 144. 二叉树的前序遍历（self solve）
// 给定二叉树的根节点, 返回其前序遍历

// 递归
func preorderTraversal(root *tree.TreeNode) []int {
	var res []int
	var build func(node *tree.TreeNode)
	build = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		build(node.Left)
		build(node.Right)
	}
	build(root)
	return res
}

// 遍历
func preorderTraversal2(root *tree.TreeNode) []int {
	var res []int
	var stack []*tree.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
