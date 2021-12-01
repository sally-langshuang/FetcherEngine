package leet

import (
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	fi, fj := 0, 0
	for j := 1; ; j++ {
		fj = fi + 9 *int(math.Pow(float64(10), float64(j - 1))) * j
		//fmt.Println(fj, "1" + strings.Repeat("0", j - 1), strings.Repeat("9", j))
		if fj < n {
			fi = fj
			continue
		}
		y := n - fi - 1
		r  := strconv.Itoa(int(math.Pow(float64(10), float64(j - 1)))+ y / j)[y % j] - '0'
		return int(r)
	}
}
