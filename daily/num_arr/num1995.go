package num_arr

//O(n4)
func countQuadruplets0(nums []int) (ans int) {
	if len(nums) < 4 {
		return 0
	}
	for l := 3; l < len(nums); l++ {
		for k := 2; k < l; k++ {
			for j := 1; j < k; j++ {
				for i := 0; i < j; i++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						ans++
					}
				}
			}
		}
	}
	return
}

// O(n2)
func countQuadruplets1(nums []int) (ans int) {
	if len(nums) < 4 {
		return 0
	}
	sumCount := map[int]int{}
	for k := len(nums) - 2; k >= 2; k-- {
		sumCount[nums[k+1]]++
		for j := k - 1; j >= 1; j-- {
			for i := j - 1; i >= 0; i-- {
				if v := nums[i] + nums[j] + nums[k]; sumCount[v] > 0 {
					ans += sumCount[v]
				}
			}
		}
	}
	return
}

// O(n2)
func countQuadruplets(nums []int) (ans int) {
	if len(nums) < 4 {
		return 0
	}
	dcCount := map[int]int{} // nums[d] - nums[c]
	for b := len(nums) - 3; b >= 1; b-- {
		for d := b + 2; d < len(nums); d++ {
			dcCount[nums[d]-nums[b+1]]++
		}
		for a := range nums[:b] {
			if v := nums[a] + nums[b]; dcCount[v] > 0 {
				ans += dcCount[v]
			}
		}

	}
	return
}
