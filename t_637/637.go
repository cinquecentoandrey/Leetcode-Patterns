package main

// https://leetcode.com/problems/average-of-levels-in-binary-tree/description/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	
	var arr []float64
	q := []*TreeNode{root}

	for len(q) > 0 {
		sum := 0.0
		l := len(q)

		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]

			if node.Right != nil {
				q = append(q, node.Right)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			sum += float64(node.Val)
		}

		arr = append(arr, sum/float64(l))
	}

	return arr
}

func main() {
	root := &TreeNode{}
	root.Val = 3

	root1 := &TreeNode{}
	root1.Val = 9

	root2 := &TreeNode{}
	root2.Val = 20

	root3 := &TreeNode{}
	root3.Val = 15

	root5 := &TreeNode{}
	root5.Val = 7

	root.Left = root1
	root.Right = root2

	root1.Left = root3
	root2.Right = root5

	fmt.Println(averageOfLevels(root))

}