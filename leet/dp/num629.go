package dp

import (
	"fmt"
	"strings"
)

func printArr(a [2][]int, cur, in, j int)  {
	fmt.Println(strings.Repeat(" ", 3 + j*2) + "V")
	for i, x:= range a{
		if cur == i {
			fmt.Println(fmt.Sprintf("->%v %d", x, in))
		}else{
			fmt.Printf("  %v\n", x)
		}

	}
	fmt.Println(strings.Repeat("-", len(a[0])))
}

func kInversePairs(n int, k int) int {
	var mod int = 1e9 + 7
	cache := [2][]int{make([]int, k + 1), make([]int, k + 1)}
	cache[1][0] = 1
	for i := 2; i <= n; i++ {
		for j:= 0; j <=k; j++ {
			cur := i & 1
			pre := cur ^ 1

			if j == 0 {
				cache[cur][j] = 1
			}else {
				cache[cur][j] = cache[cur][j - 1] + cache[pre][j]
				if i <= j {
					cache[cur][j] -= cache[pre][j-i]
				}
			}
			cache[cur][j] %= mod
			if cache[cur][j] < 0 {
				cache[cur][j] += mod
			}
		}
	}
	return cache[n & 1][k]
}
