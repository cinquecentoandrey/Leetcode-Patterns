package main

// https://leetcode.com/problems/minimum-depth-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	depth := 1
	q := []*TreeNode{root}

	for len(q) > 0 {
        l := len(q)

		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		depth++
	}

	return depth
}
