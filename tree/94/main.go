package main

import "algorithm-go/tree"

// 94、简单中序遍历，递归实现和遍历实现

func inorderTraversal1(root *tree.TreeNode) (res []int) {
	var inorder func(node *tree.TreeNode)
	inorder = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func inorderTraversal2(root *tree.TreeNode) (res []int) {
	var stack []*tree.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
