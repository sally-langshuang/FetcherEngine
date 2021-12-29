package dp

import "math/big"

func uniquePaths(m int, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(m-1)).Int64())
}
func uniquePaths0(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	f := [2][]int{}
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][1] = 1

	for i := 1; i <= m; i++ {
		i1 := i % 2
		f[i1] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			f[i1][j] = f[1-i1][j] + f[i1][j-1]
		}
	}
	return f[m%2][n]
}
