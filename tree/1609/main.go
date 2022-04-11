package main

import (
	"algorithm-go/tree"
	"math"
)

// 1609. 奇偶树

func isEvenOddTree(root *tree.TreeNode) bool {
	layer := 0
	stack := []*tree.TreeNode{root}
	for len(stack) > 0 {
		size := len(stack)
		var prev int
		if isEven(layer) {
			prev = 0
		} else {
			prev = math.MaxInt32
		}
		for i := 0; i < size; i++ {
			curr := stack[0]
			stack = stack[1:]
			if isEven(layer) {
				if prev >= curr.Val || isEven(curr.Val) {
					return false
				}
			} else {
				if prev <= curr.Val || !isEven(curr.Val) {
					return false
				}
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			prev = curr.Val
		}
		layer++
	}
	return true
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
