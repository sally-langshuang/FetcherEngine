package leet

import (
	"math"
)

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//1,1,-1,-1,3  -1
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	ans := nums[0]+nums[1] + nums[2]
	for first := 0; first < n -2; first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		for second, third := first + 1, n - 1; second < third; {
			if second > first + 1 && nums[second] == nums[second - 1] {
				second++
				continue
			}
			if third < n - 1 && nums[third] == nums[third + 1] {
				third--
				continue
			}
			a:= nums[first] + nums[second] + nums[third]
			if math.Abs(float64(a - target)) < math.Abs(float64(ans - target)) {
				ans = a
			}
			if a - target > 0 {
				third--
			} else {
				second++
			}
		}
	}
	return ans
}
