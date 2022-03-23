package main

import (
	"fmt"

	"algorithm-go/tree"
)

// 95. 不同的二叉搜索数II
func generateTrees(n int) []*tree.TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}

func build(low, high int) (res []*tree.TreeNode) {
	if low > high {
		res = append(res, nil)
		return res
	}
	// 1. 穷举root节点的所有可能
	for i := low; i <= high; i++ {
		// 2. 递归构造左右子树
		leftTree := build(low, i-1)
		rightTree := build(i+1, high)
		// 3. 穷举root节点所有子树的组合
		for _, left := range leftTree {
			for _, right := range rightTree {
				curr := tree.TreeNode{
					Val: i,
				}
				curr.Left = left
				curr.Right = right
				res = append(res, &curr)
			}
		}
	}
	return res
}

func main() {
	for _, ret := range generateTrees(3) {
		fmt.Printf("%v\n", ret)
	}
}
