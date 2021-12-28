package handle_str

/**
multiplication 31 can be replaced by a shift << 5 and a subtraction -i for better performance
31  = 2^5
i * 31 = i * (2^5 - 1) = i * 2^5 - i = i<<5 - i
*/
func longestDupSubstring(s string) string {
	f := make([]map[int][]int, len(s))
	for i := range f {
		f[i] = map[int][]int{}
	}
	for i := 1; i < len(s); i++ {
		if len(f[i-1]) != 0 {
			for l, starts := range f[i-1] {
				tmp := []int{}
				for _, start := range starts {

					if s[start+l] == s[i] {
						tmp = append(tmp, start)
					}
				}
				if len(tmp) > 0 {
					f[i][l+1] = tmp
				}
			}
		}
		tmp := []int{}
		for j := 0; j < i; j++ {
			if s[j] == s[i] {
				tmp = append(tmp, j)
			}
		}
		if len(tmp) != 0 {
			f[i][1] = tmp
		}
	}
	ans := ""
	for i := range f {
		maxL := 0
		for l, _ := range f[i] {
			if l > maxL {
				maxL = l
			}
		}
		if maxL > len(ans) {
			ans = s[f[i][maxL][0] : f[i][maxL][0]+maxL]
		}
	}
	return ans
}
