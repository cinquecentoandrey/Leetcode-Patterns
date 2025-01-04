package main

// https://leetcode.com/problems/binary-tree-paths/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// херня от андрея но работающая
func dfs(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}

	path = path + fmt.Sprint(node.Val)

	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path)
	} else {
		dfs(node.Left, path + "->", paths)
		dfs(node.Right, path + "->", paths)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string

	if root != nil {
		dfs(root, "", &paths)
	}

	return paths
}


// херня от андрея с глобальной переменной работать не будет 
// var paths []string

// func backTarcking(node *TreeNode, path string) {
// 	if node == nil {
// 		return
// 	}
// 	path = path + fmt.Sprint(node.Val)
// 	if node.Left == nil && node.Right == nil {
// 		paths = append(paths, path)
// 	} else {
// 		backTarcking(node.Left, path  + "->")
// 		backTarcking(node.Right, path  + "->")
// 	}
	
// }

// func binaryTreePaths(root *TreeNode) []string {
// 	if root != nil {
// 		path := fmt.Sprint(root.Val)
// 		backTarcking(root.Left, path  + "->")
// 		backTarcking(root.Right, path  + "->")
// 		return paths
// 	}
// 	return paths
// }