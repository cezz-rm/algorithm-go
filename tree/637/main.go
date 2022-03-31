package main

import (
	"algorithm-go/tree"
)

// 637. 二叉树的层平均值

func averageOfLevels(root *tree.TreeNode) []float64 {
	var res []float64
	stack := []*tree.TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		num, count := 0, 0
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			num += node.Val
			count += 1
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res = append(res, float64(num)/float64(count))
	}
	return res
}
