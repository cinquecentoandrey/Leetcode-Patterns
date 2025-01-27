package main

import "fmt"

// https://leetcode.com/problems/maximum-product-subarray/description/

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}

	res := maxArr(nums)
	maxVal := 1
	minVal := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			maxVal = 1 
			minVal = 1
			continue
		}
		tmp := maxVal * nums[i]
		maxVal = max(nums[i] * maxVal, nums[i] * minVal, nums[i])
		minVal = min(tmp, nums[i] * minVal, nums[i])
		res  = max(res, maxVal)
	}

	return res
}

func maxArr(nums []int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return res
}


func main() {
	//fmt.Println(maxProduct([]int{2,3,-2,4}))
	
	fmt.Println(maxProduct([]int{-4,-3,-2}))
}

// func maxProduct(nums []int) int {
//     curProd := 0
// 	prevProd := 0
// 	maxProd := nums[0]
// 	//positiveProd := 0
// 	negativeProd := 0
// 	isSubarray := true
// 	for i := 0; i < len(nums); i++ {
// 		prevProd = nums[i]
// 		curProd = 0
// 		curMaxProd := 0
// 		if i == len(nums) - 1 {
// 			if nums[i] > maxProd {
// 				maxProd = nums[i]
// 			}
// 			break
// 		}
// 		for j := i+1; j < len(nums); j++ {
// 			fmt.Println("prevProd: ", prevProd)
// 			curProd = prevProd * nums[j]
// 			if nums[j] == 0 || curProd == 0{
// 				break
// 			}
			
// 			fmt.Println("i: ", i)
// 			fmt.Println("j: ", j)
// 			fmt.Println("curProd: ", curProd)
// 			fmt.Println("nums[j]: ", nums[j])

// 			if isSubarray {
// 				if nums[j] > 0 && curProd > 0 {
// 					curMaxProd = max(curProd, prevProd)
// 				} else if nums[j] < 0  && curProd < 0 { 
// 					negativeProd = curMaxProd
// 					isSubarray = false
// 				} else if nums[j] < 0  && curProd > 0 {
// 					curMaxProd = max(curProd, prevProd)
// 				}
// 			} else {
// 				if nums[j] > 0 && curProd < 0 {
// 					negativeProd = curProd
// 				} else if nums[j] < 0 && curProd > 0 {
// 					curMaxProd = max(curProd, negativeProd)
// 					isSubarray = true
// 				}
// 			}
			
// 			// fmt.Println("curMaxProd: ", curMaxProd)
// 			// fmt.Println("posiProd: ", positiveProd)
// 			// fmt.Println()
// 		}
// 		fmt.Println("maxProd: ", maxProd)
// 		fmt.Println("curMaxProd: ", curMaxProd)
// 		maxProd = max(maxProd, curMaxProd)
// 		fmt.Println()
// 	}

// 	return maxProd
// }
