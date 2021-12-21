package num_arr

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	ans := 0
	sort.Ints(heaters)
	sort.Ints(houses)
	for i, j := 0, 0; i < len(houses) && j < len(heaters); i++ {
		l := int(math.Abs(float64(houses[i] - heaters[j])))
		for ; j+1 < len(heaters); j++ {
			v := int(math.Abs(float64(houses[i] - heaters[j+1])))
			if v > l {
				break
			}
			l = v
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
