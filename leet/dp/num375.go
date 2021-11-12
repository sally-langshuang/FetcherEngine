package dp

import (
	"fmt"
	"math"
)

func getMoneyAmount(n int) int {
	f := makeIntArr(n+1, n+1)
	//  上半部分 右下脚开始 按行
	//for i := n - 1; i >= 1; i-- {
	//	for j := i + 1; j <= n; j++ {
	//		minCost := math.MaxInt32
	//		for k := i; k < j; k++ {
	//			cost := k + max(f[i][k-1], f[k+1][j])
	//			if cost < minCost {
	//				minCost = cost
	//			}
	//		}
	//		f[i][j] = minCost
	//		fmt.Println(ArrString(f, &point{i, j}))
	//	}
	//}
	// 上半部分 左上角开始 按列
	for j := 1; j <= n; j++ {
		for i := j - 1; i >= 1; i-- {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
			fmt.Println(ArrString(f, &point{i, j}))
		}
	}
	fmt.Println(ArrString(f, &point{1,n}))
	return f[1][n]

}