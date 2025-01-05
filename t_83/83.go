package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// ничего сложного, два указателя и сдвигаем их по условию
func deleteDuplicatesListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := head.Next

	for right != nil {
		if left.Val == right.Val {
			left.Next = right.Next
			right = right.Next
		} else {
			left = right
			right = right.Next
		}
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	list1 := &ListNode{}
	list1.Val = 1

	list2 := &ListNode{}
	list2.Val = 1

	list3 := &ListNode{}
	list3.Val = 2

	list4 := &ListNode{}
	list4.Val = 3

	list5 := &ListNode{}
	list5.Val = 3

	list1.Next = list2
	list2.Next = list3
	list3.Next = list4
	list4.Next = list5

	printList(list1)
	list1 = deleteDuplicatesListNode(list1)
	fmt.Println()
	printList(list1)

	// /

}