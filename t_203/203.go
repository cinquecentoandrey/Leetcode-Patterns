package main

// https://leetcode.com/problems/remove-linked-list-elements/description/

type ListNode struct {
    Val int
    Next *ListNode
}

// просто меняем указатель ноды, и про доп ноду в начале не забываем
func removeElements(head *ListNode, val int) *ListNode {
	temp := &ListNode{}
	temp.Val = -1
	temp.Next = head

	cur := temp

	for cur.Next != nil {
		if (cur.Next.Val == val) {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return temp.Next
}