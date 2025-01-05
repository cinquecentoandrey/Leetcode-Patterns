package main

// https://leetcode.com/problems/merge-two-sorted-lists/description/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// просто сраниваем значения и двигаем ноды
func mergeTwoSortLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	cur := temp

	for list1 != nil && list2 != nil {
		if (list1.Val >= list2.Val) {
			cur.Next = list2
			list2 = list2.Next
		} else  {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return temp.Next
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

	list1_1 := &ListNode{}
	list1_1.Val = 2

	list1_2 := &ListNode{}
	list1_2.Val = 4

	list1.Next = list1_1
	list1_1.Next = list1_2

	list2 := &ListNode{}
	list2.Val = 1

	list2_1 := &ListNode{}
	list2_1.Val = 3

	list2_2 := &ListNode{}
	list2_2.Val = 4

	list2.Next = list2_1
	list2_1.Next = list2_2

	merged := mergeTwoSortLists(list1, list2)
	printList(merged)

}