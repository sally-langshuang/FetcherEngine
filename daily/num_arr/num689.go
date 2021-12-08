package num_arr

func sumKPeriod(starts []int, nums[]int, k int) int{
	var ans int
	for x:= 0; x < k; x++ {
		for _, start := range starts {
			ans += nums[start + x]
		}
	}
	return ans
}
func maxSumOfThreeSubarrays0(nums []int, k int) []int {
	f := make([][][]int, k + 1)
	for j:=0; j < k + 1; j++ {
		f[j] = make([][]int, 3)
	}
	for j :=0; j < len(nums); j++ {
		jj := j % (k + 1)
		j1 := (j - 1) % (k + 1)
		jk:= (j - k) % (k + 1)
		f[jj] = make([][]int, 3)
		for i:=0; i < 3; i++ {
			if j < k * (i + 1) - 1 {continue}
			f[jj][i] = []int{}
			if i > 0 {
				f[jj][i] = append(f[jj][i], f[jk][i - 1]...)
			}
			f[jj][i] = append(f[jj][i], j- k + 1)
			ans1 := sumKPeriod(f[jj][i], nums, k )
			if j- 1 >= 0 && f[j1][i] != nil {
				ans2 := sumKPeriod(f[j1][i], nums, k)
				if ans2 >= ans1 {
					f[jj][i] = f[j1][i]
				}
			}
		}
	}
	return f[(len(nums) - 1) % (k + 1)][2]
}
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n:= len(nums)
	var sum1, sum2, sum3, maxSum1, maxSum2, maxSum3, maxSum1Idx, maxSum12idx1, maxSum12idx2 int
	var ans []int
	for i:= 0; i < n - 2 * k; i++ {
		sum1 += nums[i]
		sum2 += nums[i + k]
		sum3 += nums[i + 2 * k]
		if i >= k - 1 {
			if sum1 > maxSum1 {
				maxSum1Idx = i - k + 1
				maxSum1 = sum1
			}
			if sum2 + maxSum1 > maxSum2 {
				maxSum12idx1, maxSum12idx2 = maxSum1Idx, i + 1
				maxSum2 = sum2 + maxSum1
			}
			if sum3 + maxSum2 > maxSum3 {
				ans = []int{maxSum12idx1, maxSum12idx2, i + k + 1}
				maxSum3 = sum3 + maxSum2
			}
			sum1 -= nums[i - k + 1]
			sum2 -= nums[i + 1]
			sum3 -= nums[i + k + 1]
		}

	}
	return ans
}
