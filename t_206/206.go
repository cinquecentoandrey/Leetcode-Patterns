package main

// https://leetcode.com/problems/reverse-linked-list/
// https://www.geeksforgeeks.org/reverse-a-linked-list/

func main() {

}

type ListNode struct {
    Val int
    Next *ListNode
}

// просто не запутаться в указателях
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