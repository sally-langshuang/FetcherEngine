package leet

func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{'(': ')', '[':']', '{': '}'}
	for i := range s {
		if c := s[i]; c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) != 0 && m[stack[len(stack) - 1]] == c {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
