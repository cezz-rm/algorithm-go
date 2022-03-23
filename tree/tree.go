package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) New(n int) TreeNode {
	return TreeNode{
		Val:   n,
		Left:  nil,
		Right: nil,
	}
}

func (t *TreeNode) String() string {
	// 关于带指针的结构体打印
	// https://www.jianshu.com/p/1b04d75769a7
	return fmt.Sprintf("{val: %v, left: %v, right: %v}", t.Val, t.Left, t.Right)
}
