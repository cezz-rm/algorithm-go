package main

import "algorithm-go/tree"

// 865.具有所有最深节点的最小子树

func subtreeWithAllDeepest(root *tree.TreeNode) *tree.TreeNode {
	var subTree func(root *tree.TreeNode) (node *tree.TreeNode, depth int)
	subTree = func(root *tree.TreeNode) (node *tree.TreeNode, depth int) {
		if root == nil {
			return nil, 0
		}
		left, leftDepth := subTree(root.Left)
		right, rightDepth := subTree(root.Right)
		if leftDepth == rightDepth {
			return root, leftDepth + 1
		}
		// 如果不相等, 返回左右节点中较深的一个
		if leftDepth > rightDepth {
			return left, leftDepth + 1
		}
		return right, rightDepth + 1
	}
	ret, _ := subTree(root)
	return ret
}
