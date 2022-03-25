package main

import (
	"algorithm-go/tree"
)

// 111.二叉树的最小深度
// 给定一个二叉树, 找出其最小深度

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return 1 + leftMinDepth + rightMinDepth
	}
	return 1 + min(leftMinDepth, rightMinDepth)
}

//func minDepth2(root *tree.TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	minD := math.MaxInt32
//	if root.Left != nil {
//		minD = min(minDepth(root.Left), minD)
//	}
//	if root.Right != nil {
//		minD = min(minDepth(root.Right), minD)
//	}
//	return 1 + minD
//}

func minDepth3(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*tree.TreeNode{root}
	depth := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
