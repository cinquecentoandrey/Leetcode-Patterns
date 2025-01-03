package main

// https://leetcode.com/problems/majority-element/description/

import "fmt"

func majoriyElement(nums []int) int {
	freqMap := make(map[int]int)
	mean := len(nums) / 2

	for _, v := range nums {
		freqMap[v] += 1
	}

	for k, v := range freqMap {
		if v > mean {
			return k
		}
	}

	return nums[0]
}

// шайтан вариант
func majoriyElement1(nums []int) int {
	var result int
	vote := 0

	for _, v := range nums {
		if vote == 0 {
			result = v
		} else if (v == result) {
			vote++
		} else {
			vote--
		}
	}

	return result
}


func main() {
	arr := []int{2,2,1,1,1,2,2}
	fmt.Println(majoriyElement(arr))

	arr1 := []int{3,2,3}

	fmt.Println(majoriyElement(arr1))
}