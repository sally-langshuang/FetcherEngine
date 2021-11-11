package leet

func makeMatrix(m, n int) [][]bool {
	res := make([][]bool, m)
	for i := 0 ; i< m; i++ {
		res[i] = make([]bool, n)
	}
	return res
}

func isMatch(s string, p string) bool {
	arr := makeMatrix(len(s) + 1, len(p) + 1)
	arr[0][0] = true
	// i 为 s 的长度 对应s[i-1]元素 ，以此类推j p
	var i, j int
	for ; i < len(s) + 1; i++ {
		for j = 1; j<len(p) + 1; j++ {
			if p[j - 1] == '*' {
				// *==0 || *==+（s不为空&& p[:j]能匹配s[:i-1], .与s[i]相同）
				//arr[i][j] = arr[i][j - 2] || (i > 0 &&(arr[i - 1][j] && (s[i - 1] ==p[j - 2] || p[j - 2] == '.')))
				arr[i][j] = i > 0 &&(arr[i - 1][j] && (s[i - 1] ==p[j - 2] || p[j - 2] == '.')) || arr[i][j - 2]
			}else if i > 0{
				arr[i][j] = arr[i - 1][j - 1] && (s[i - 1] == p[j - 1] || p[j - 1] == '.')
			}
		}
	}
	return arr[i-1][j-1]
}
