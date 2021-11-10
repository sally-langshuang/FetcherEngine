package leet
func isMatch(s string, p string) bool {
	var i int
	var lastChar uint8
	for _, r := range p {
		if 'a' <= r && r <= 'z' {
			if string(s[i]) != string(r) {
				return false
			} else {
				lastChar = s[i]
				i++
			}
		} else if r == '.' {
			lastChar = s[i]
			i++
		} else if r == '*' {
			if s[i] != lastChar {
				return false
			} else {
				i++
			}
		}
	}
	return true
}
