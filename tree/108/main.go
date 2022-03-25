package main

import "algorithm-go/tree"

// 108. 将有序数组转换为二叉搜索数

func sortedArrayToBST(nums []int) *tree.TreeNode {
	return build(nums, 0, len(nums)-1)
}

// 将闭区间[left, right]中的元素转化成BST
func build(nums []int, left, right int) *tree.TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := tree.TreeNode{
		Val: nums[mid],
	}
	// 构建左子树
	root.Left = build(nums, left, mid-1)
	// 构建右子树
	root.Right = build(nums, mid+1, right)
	return &root
}
