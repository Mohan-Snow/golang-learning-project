package tests

import (
	"golang-learning-project/leetcode"
	"testing"
)

func BenchmarkInvertTree(b *testing.B) {
	leftNode := leetcode.TreeNode{Val: 1, Left: nil, Right: nil}
	rightNode := leetcode.TreeNode{Val: 3, Left: nil, Right: nil}
	root := leetcode.TreeNode{Val: 2, Left: &leftNode, Right: &rightNode}
	for i := 0; i < b.N; i++ {
		leetcode.InvertTree(&root)
	}
}
