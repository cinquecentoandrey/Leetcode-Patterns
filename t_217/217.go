package t217

// https://leetcode.com/problems/contains-duplicate/description/

func containsDuplicate(nums []int) bool {
	myMap := map[int]bool{}
	for _, v := range nums {
		if myMap[v] {
			return true
		} 
		myMap[v] = true 
	}

	return false
}