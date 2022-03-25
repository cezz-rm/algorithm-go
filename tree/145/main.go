package main

import "algorithm-go/tree"

// 145. 二叉树的后续遍历

// 递归
func postorderTraversal(root *tree.TreeNode) []int {
	var res []int
	var build func(node *tree.TreeNode)
	build = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		build(node.Left)
		build(node.Right)
		res = append(res, node.Val)
	}
	build(root)
	return res
}

// 遍历
func postorderTraversal2(root *tree.TreeNode) []int {
	var res []int
	var stack []*tree.TreeNode
	var prev *tree.TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			res = append(res, node.Val)
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return res
}
