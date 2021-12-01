package handle_str

func maxPower(s string) int {
	maxL, l := 1, 1
	for i:= 1; i < len(s); i++{
		if s[i] != s[i - 1] {
			l = 1
		} else {
			l++
			if l > maxL {
				maxL = l
			}
		}
	}
	return maxL
}
