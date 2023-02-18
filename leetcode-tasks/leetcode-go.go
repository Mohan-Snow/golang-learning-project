package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	leftNode := TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode := TreeNode{Val: 3, Left: nil, Right: nil}
	root := TreeNode{Val: 2, Left: &leftNode, Right: &rightNode}
	fmt.Printf("%v", root)
	invertTree(&root)
	fmt.Printf("%v", root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
