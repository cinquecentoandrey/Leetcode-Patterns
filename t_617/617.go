package main

// https://leetcode.com/problems/merge-two-binary-trees/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	mergeTrees(root1.Left, root2.Left)
	mergeTrees(root1.Right, root2.Right)

	return root1
}