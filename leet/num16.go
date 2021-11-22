package leet

import (
	"math"
	"sort"
)

//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//1,1,-1,-1,3  -1
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j, k := i + 1, len(nums) - 1; j < k; {
			if j > i + 1 && nums[j - 1] == nums[j] {
				j++
				continue
			}
			if k < len(nums) - 1 && nums[k] == nums[k + 1] {
				k--
				continue
			}
			x := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(x - target)) < math.Abs(float64(closest - target)) {
				closest = x
			}
			if x - target > 0 {
				k--
			} else {
				j++
			}

		}
	}
	return closest
}
