package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	current := tmp
	carry := 0

	var sum int
	for l1 != nil || l2 != nil || carry > 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return tmp.Next
}

// несложно, много повторяющегося кода, рефакторить мне лень
// андрюша накатал простыню
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 == nil {
//		return nil
//	}
//
//	res := &ListNode{}
//	curRes := res
//
//	var sumVal int
//	var div int
//	//var remainder int
//	for l1 != nil && l2 != nil {
//		sumVal = l1.Val + l2.Val + div
//		if sumVal >= 10 {
//			div = sumVal / 10
//			sumVal = sumVal % 10
//		} else {
//			div = 0
//		}
//		curNode := &ListNode{Val: sumVal}
//		curRes.Next = curNode
//		curRes = curNode
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//
//	if l1 == nil && l2 != nil {
//		for l2 != nil {
//			sumVal = l2.Val + div
//			if sumVal >= 10 {
//				div = sumVal / 10
//				sumVal = sumVal % 10
//			} else {
//				div = 0
//			}
//			curNode := &ListNode{Val: sumVal}
//			curRes.Next = curNode
//			curRes = curNode
//			l2 = l2.Next
//		}
//	} else if l2 == nil && l1 != nil {
//		for l1 != nil {
//			sumVal = l1.Val + div
//			if sumVal >= 10 {
//				div = sumVal / 10
//				sumVal = sumVal % 10
//			} else {
//				div = 0
//			}
//			curNode := &ListNode{Val: sumVal}
//			curRes.Next = curNode
//			curRes = curNode
//			l1 = l1.Next
//		}
//	}
//
//	if div != 0 {
//		curRes.Next = &ListNode{Val: div}
//	}
//
//	return res.Next
//}

func main() {
	l1 := createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createLinkedList([]int{9, 9, 9, 9})

	printLinkedList(addTwoNumbers(l1, l2))
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
}
