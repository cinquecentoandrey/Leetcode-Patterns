package main

// https://leetcode.com/problems/find-the-duplicate-number/description/
// https://www.geeksforgeeks.org/floyds-cycle-finding-algorithm/
// https://www.geeksforgeeks.org/find-any-one-of-the-multiple-repeating-elements-in-read-only-array-set-2/

// для умников, сам по условию o(n)/o(1) не осилил (но до фишки с циклом додумался), с мапой сортировкой и прочим смог бы
// 
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if fast == slow {
			break
		}
	}

	slow = nums[0]

	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}