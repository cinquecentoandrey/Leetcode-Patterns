package main

// https://leetcode.com/problems/counting-bits/description/
// дада сдвигаем итд

import "fmt"

func main() {
	fmt.Println(countBits(5))
	fmt.Println(5 >> 1)
}

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		res[i] = count(i)
	}

	return res
}

func count(n int) int {
	var cnt int
	for {
		if n == 0 {
			break
		}
		div := n % 2
		if div == 1 {
			cnt++
		}
		n = n / 2
	}
	return cnt
}