package main

// https://leetcode.com/problems/middle-of-the-linked-list/

func main() {

}

type ListNode struct {
    Val int
    Next *ListNode
}


// опять же 2 указателя и пошли
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