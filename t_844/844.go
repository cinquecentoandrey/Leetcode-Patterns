package main

// https://leetcode.com/problems/backspace-string-compare/


// идея в том, что идем с конца строки и двигаем индекс в заивимости от символа и кол-ва решеток
// пишут, что Easy, но как по мне это medium
func getCurrentIndex(str string, idx int) int {
	toSkip := 0

	for idx >= 0 {
		if str[idx] == '#' {
			toSkip++
		} else if toSkip > 0 {
			toSkip--
		} else {
			break
		}
		idx--
	}

	return idx
}


func backspaceCompare(s string, t string) bool {
    i := len(s) - 1
	j := len(t) - 1

	for i >= 0 || j >= 0 {
		i = getCurrentIndex(s, i)
		j = getCurrentIndex(t, j)

		if i < 0 && j < 0 {
			return true
		} 
		if i < 0 || j < 0 {
			return false
		} 
		if s[i] != t[j] {
			return false
		}
        i--
        j--
	}

	return true
}

func main() {
	
}