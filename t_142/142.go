package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// trotoise
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}

	slow2 := head

	for slow.Next != nil {
		if slow == slow2 {
			return slow
		} 
		slow = slow.Next
		slow2 = slow2.Next
	}

	return nil
}

// o(n)
// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}

// 	nodeMap := make(map[*ListNode]bool)

// 	for head.Next != nil {
// 		_, exists := nodeMap[head]
// 		if exists {
// 			return head
// 		} else {
// 			nodeMap[head] = true
// 			head = head.Next
// 		}
// 	}

//     return nil
// }

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(detectCycle(node1))
}

// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	currentIdx := 0
// 	nodeMap := make(map[NodeKey]int)
// 	tmpHead := head
// 	var doubleKey NodeKey
// 	for {
// 		fmt.Println(1)
// 		curNodeKey := NodeKey{currentIdx, head.Val}

// 		_, exists := nodeMap[curNodeKey]
// 		if exists {
// 			fmt.Println()
// 			doubleKey = NodeKey{currentIdx, head.Val}
// 			break
// 		} else {
// 			nodeMap[curNodeKey] = 1
// 			head = head.Next
// 			currentIdx++
// 		}
// 	}

// 	currentIdx = 0
// 	for {
// 		if currentIdx == doubleKey.nodePos && head.Val == doubleKey.nodeVal {
// 			return tmpHead
// 		} else {
// 			tmpHead = tmpHead.Next
// 			currentIdx++
// 		}
// 	}

// 	return nil
// }

