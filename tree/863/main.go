package main

import "algorithm-go/tree"

// 863. 二叉树中所有距离为k的结点

func distanceK(root, target *tree.TreeNode, k int) []int {
	var ans []int
	parents := map[int]*tree.TreeNode{}
	var findParents func(node *tree.TreeNode)
	findParents = func(node *tree.TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)

	var findAns func(node, from *tree.TreeNode, depth int)
	findAns = func(node, from *tree.TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			ans = append(ans, node.Val)
			return
		}
		if node.Left != from {
			findAns(node.Left, node, depth+1)
		}
		if node.Right != from {
			findAns(node.Right, node, depth+1)
		}
		if parents[node.Val] != from {
			findAns(parents[node.Val], node, depth+1)
		}
	}
	findAns(target, nil, 0)
	return ans
}
