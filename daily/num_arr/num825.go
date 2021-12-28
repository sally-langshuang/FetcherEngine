package num_arr

import (
	"math"
	"sort"
)

func numFriendRequests(ages []int) int {
	ans := 0
	sort.Ints(ages)
	ages = ages[sort.SearchInts(ages, 15):]
	for i, left, right := len(ages)-1, len(ages)-1, len(ages)-1; i >= 0; i-- {
		minY, _ := math.Modf(0.5*float64(ages[i]) + 7)
		minY++
		//left := sort.SearchInts(ages, int(minY))
		for left > 0 && ages[left-1] >= int(minY) {
			left--
		}
		for right > 0 && ages[right] > ages[i] {
			right--
		}
		//right := sort.SearchInts(ages, ages[i]+1)
		//ans += right - left - 1
		ans += right - left
	}
	return ans
}
