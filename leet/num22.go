package leet

import "strings"

func countParenthesisNum(n int) int{
	if n <= 0 {
		return 1
	}
	f := makeMatrixInt(n + 1, n + 1)
	for j := 0; j <= n; j++ {
		for i := 0; i <= n; i++ {
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

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	f := makeMatrixStringArr(n + 1, n + 1)
	for j := 0; j <= n; j++ {
		for i := 0; i <= n; i++ {
			if j == 0 {
				f[i][j] = []string{strings.Repeat(")", i)}
			} else {
				f[i][j] = AddString(nil, "(" , f[i+1][j-1])
				if i > 0 {
					f[i][j] = append(f[i][j], AddString(nil, ")", f[i-1][j])...)
				}
			}
		}
	}
	return f[0][n]

}
