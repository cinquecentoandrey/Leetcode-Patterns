package main

// https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)

	return root
}