package main

// https://leetcode.com/problems/palindrome-linked-list/

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
	   return head
   }

   slow := head
   fast := head.Next

   for fast.Next != nil && fast.Next.Next != nil {
	   slow = slow.Next
	   fast = fast.Next.Next
   }

   return slow.Next
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var prev *ListNode

	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
        prev = cur
        cur = nxt
	}

	return prev
}   

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	middle := middleNode(head)
	revMiddle := reverseList(middle)

	for revMiddle != nil {
		if head.Val != revMiddle.Val {
			return false
		}
		head = head.Next
		revMiddle = revMiddle.Next
	}
	

	return true
}