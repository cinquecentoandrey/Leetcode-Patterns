package main

import "fmt"

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func traverse(k int, digits string, path string, paths *[]string, phoneKeyMap map[int][]string) {

	if k == len(digits) {
		*paths = append(*paths, path)
		return
	}
	
	curNum := int(digits[k] - '0')
	
	for i := 0; i < len(phoneKeyMap[curNum]); i++ {
		newPath := path + phoneKeyMap[curNum][i]
		traverse(k+1, digits, newPath, paths, phoneKeyMap)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var paths []string
	phoneKeyMap := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	if len(digits) == 1 {
		return phoneKeyMap[int(digits[0] - '0')]
	}
	
	traverse(0, digits, "", &paths, phoneKeyMap)
	return paths
}

func main() {
	fmt.Println(letterCombinations("23"))
}

