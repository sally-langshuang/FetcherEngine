package tree

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i, j := 0, len(s) - 1; i <= j; {
		if (s[i] != s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
