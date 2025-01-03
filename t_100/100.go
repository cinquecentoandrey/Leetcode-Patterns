package main

// https://leetcode.com/problems/same-tree/

// dfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	
	if p == nil || q == nil {
		return false
	}
    
    if p.Val != q.Val {
        return false
    }

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


// bfs
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}
// 	if p == nil || q == nil {
// 		return false
// 	}
    
// 	pq := []*TreeNode{p}
// 	qq :=  []*TreeNode{q}

// 	for len(pq) > 0 || len(qq) > 0 {
// 		pqLen := len(pq)
// 		qqLen := len(qq)

// 		if pqLen != qqLen {
// 			return false
// 		}

// 		for i := 0; i < pqLen; i++ {
// 			pqNode := pq[0]
// 			pq = pq[1:]

// 			qqNode := qq[0]
// 			qq = qq[1:]

// 			if pqNode.Val != qqNode.Val {
// 				return false
// 			}

// 			if pqNode.Left != nil && qqNode.Left != nil {
// 				pq = append(pq, pqNode.Left)
// 				qq = append(qq, qqNode.Left)
// 			}

// 			if pqNode.Right != nil && qqNode.Right != nil {
// 				pq = append(pq, pqNode.Right)
// 				qq = append(qq, qqNode.Right)
// 			}

// 			if (pqNode.Left != nil && qqNode.Left == nil) || (pqNode.Left == nil && qqNode.Left != nil) {
// 				return false
// 			}

// 			if (pqNode.Right != nil && qqNode.Right == nil) || (pqNode.Right == nil && qqNode.Right != nil) {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }