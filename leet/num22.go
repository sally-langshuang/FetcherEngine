package leet

import (
	"strings"
)

func countParenthesisNum(n int) int{
	if n <= 0 {
		return 1
	}
	f := makeMatrixInt(n + 1, n + 1)
	for j := 0; j <= n; j++ {
		for i := 0; i <= n - j; i++ {
			if j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i+1][j-1]
				if i > 0 {
					f[i][j] += f[i-1][j]
				}
			}
		}
	}
	return f[0][n]
}

func AddString(preffix []string, e string, subffix []string) []string {
	if preffix == nil ||  len(preffix) == 0 {
		preffix = []string{""}
	}
	if subffix == nil ||  len(subffix) == 0 {
		subffix = []string{""}
	}
    ans := make([]string, 0)
	for _, p := range preffix {
		for _, s := range subffix {
			ans = append(ans, p + e + s)
		}

	}
	return ans
}
func f(i, l, r int, s []uint8, ans *[]string) {
	if i == len(s) {
		if l == r {
			*ans = append(*ans, string(s))
		}
	} else {
		s[i] = '('
		if l + 1 >= r && l + 1 <= len(s) / 2 {
			f(i + 1, l+1, r, s, ans)
		}
		s[i] = ')'
		if l >= r + 1 {
			f(i + 1, l, r+ 1, s, ans)
		}

	}
}
func generateParenthesis2(n int) []string {
	s := make([]uint8, 2*n)
	ans := make([]string, 0)
	f(0,0, 0, s, &ans)
	return ans
}

func generateParenthesis3(n int) []string {
	if n <= 0 {
		return []string{}
	}
	f := makeMatrixStringArr(n + 1, 2)
	for j := 0; j <= n; j++ {
		realj := j % 2
		prej := realj ^ 1
		for i := 0; i <= n-j; i++ {
			if j == 0 {
				f[i][realj] = []string{strings.Repeat(")", i)}
			} else {
				f[i][realj] = AddString(nil, "(" , f[i+1][prej])
				if i > 0 {
					f[i][realj] = append(f[i][realj], AddString(nil, ")", f[i-1][realj])...)
				}
			}
		}
	}
	return f[0][n%2]

}

func generateParenthesisf(i, l, r int, s []uint8, ans *[]string) {
	if i == len(s) {
		if l == r {
			*ans = append(*ans, string(s))
		}
	} else {
		s[i] = '('
		if l + 1 >= r  && l + 1 <= len(s) / 2 {
			generateParenthesisf(i + 1, l+1, r, s, ans)
		}
		s[i] = ')'
		if l >= r + 1 {
			generateParenthesisf(i + 1, l, r+ 1, s, ans)
		}

	}
}
func generateParenthesis(n int) []string {
	s := make([]uint8, 2*n)
	ans := make([]string, 0)
	generateParenthesisf(0,0, 0, s, &ans)
	return ans
}
