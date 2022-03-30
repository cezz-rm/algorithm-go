package main

import (
	"algorithm-go/tree"
	"fmt"
	"strconv"
)

// 606. 根据二叉树创建字符串

func tree2str(root *tree.TreeNode) string {
	if root == nil {
		return ""
	}
	left := tree2str(root.Left)
	right := tree2str(root.Right)
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	if root.Left != nil && root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + left + ")"
	}
	if root.Left == nil && root.Right != nil {
		return strconv.Itoa(root.Val) + "()" + "(" + right + ")"
	}
	return strconv.Itoa(root.Val) + "(" + left + ")" + "(" + right + ")"
}

// 官方
func tree2str2(root *tree.TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
