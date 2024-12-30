package main

// https://leetcode.com/problems/range-sum-query-immutable/
// дада префиксная сумма, но я тупой

type NumArray struct {
    arr []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
		arr: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int
    for i := left; i <= right; i++ {
		sum += this.arr[i]
	}
	return sum
}