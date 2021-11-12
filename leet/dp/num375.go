package dp

import (
	"fmt"
	"math"
)

func getMoneyAmount(n int) int {
	f := makeIntArr(n+1, n+1)
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}
	fmt.Println(ArrString(f))
	return f[1][n]

}