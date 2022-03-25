package main

import "algorithm-go/tree"

// 501. 查找二叉搜索数的众数
// 出现次数最多的数

func findMode(root *tree.TreeNode) []int {
	var res []int
	var prev, curCount, maxCount int

	update := func(x int) {
		if x == prev {
			curCount++
		} else {
			prev, curCount = x, 1
		}
		if curCount == maxCount {
			res = append(res, prev)
		} else if curCount > maxCount {
			maxCount = curCount
			res = []int{prev}
		}
	}

	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}
