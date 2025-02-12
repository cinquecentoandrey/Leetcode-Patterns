package main

// https://leetcode.com/problems/palindromic-substrings/

// была уже подобная https://leetcode.com/problems/longest-palindromic-substring/description/
func countSubstrings(s string) int {
    res := 0 

	for i := 0; i < len(s); i++ {
		l := i
		r := i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
            r++
		}

		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
            r++
		}
	}

	return res
}