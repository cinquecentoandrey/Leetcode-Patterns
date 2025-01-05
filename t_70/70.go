package main

// https://leetcode.com/problems/climbing-stairs/description/

import "fmt"

func main() {
	fmt.Println(climbStairs(6))
}

// фибо
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	first, second := 1, 2

	for i := 0; i < n-2; i++ {
		first, second = second, first + second
	}

	return second
}