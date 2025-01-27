package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring/description/

func search(r int, l int, cur int, s string) string{
	tmpRes := ""
	
	for {
		if l >= 0 && r < len(s) && s[l] == s[r]  {
			tmpRes = string(s[l]) + string(s[cur]) + string(s[r])
			// fmt.Println(tmpRes)
			fmt.Printf("l: %d, r: %d, cur: %d\n",l,r,cur)
			r--
			l++
		} else if l < 0 && s[r] == s[cur] {
			tmpRes = string(s[cur]) + string(s[r])
			r--
		} else if  r > len(s) && s[l] == s[cur] {
			tmpRes = string(s[l]) + string(s[cur])
			l++
		} else {
			break
		}
	} 

	return tmpRes
}

func longestPalindrome(s string) string {
    res := ""
	resLen := 0
	l := 0
	r := 0
	for i := 0; i < len(s); i++ {
		l = i
		r = i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > resLen {
				res = s[l:r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}

		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r - l + 1 > resLen {
				res = s[l:r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}
	}

	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}