package main

// https://leetcode.com/problems/diameter-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lp(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := lp(root.Left, diameter)
	right := lp(root.Right, diameter)
    
	*diameter = max(*diameter, left+right)

	return 1 + max(left, right)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := 0
	lp(root, &diameter)
	return diameter
}

