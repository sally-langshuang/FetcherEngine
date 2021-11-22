package leet

import "sort"

func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	ans = make([][]int, 0)
	for i := 0; i < len(nums) - 3; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < len(nums) - 2; j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			for k, m := j + 1, len(nums) - 1; k < len(nums) - 1 && k < m;  {
				if k > j + 1 && nums[k] == nums[k - 1] {
					k++
					continue
				}
				if m < len(nums) - 1 && nums[m] == nums[m + 1] {
					m--
					continue
				}
				if x := nums[i] + nums[j] + nums[k] + nums[m] - target; x > 0 {
					m--
				} else if x < 0 {
					k++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[m]})
					m--
					k++
				}
			}
		}
	}
	return ans
}
