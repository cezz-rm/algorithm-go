package main

import "algorithm-go/tree"

// 701.二叉搜索树的插入操作

func insertIntoBST(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func insertIntoBST2(root *tree.TreeNode, val int) *tree.TreeNode {
	if root == nil {
		return &tree.TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &tree.TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &tree.TreeNode{Val: val}
				break
			}
			p = p.Right
		}
	}
	return root
}
