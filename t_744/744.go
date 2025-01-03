package main

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

func nextGreaterLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	mid := 0

	for left <= right {
		mid = left + (right - left) / 2

		if target < letters[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return letters[left % len(letters)]
}