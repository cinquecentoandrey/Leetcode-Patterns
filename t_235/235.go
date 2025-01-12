package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node := root

	for node != nil {
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			return node
		}
	}

	return nil
 }