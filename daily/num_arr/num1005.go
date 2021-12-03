package num_arr

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	nums = heakSort(nums)
	return largest(nums, k)
	
}

func largest(nums []int, k int) int{
	if len(nums) == 0 {
		return 0
	}
	heakInit(nums)
	count, m, l, lastNegAbs, firstPosAbs := 0, 0, len(nums), -1, -1
	for i := 0; i < l; i++ {
		v := pop(&nums)
		if v < 0 && i < k {
			m++
			count -= v
			lastNegAbs = -v
		} else {
			if firstPosAbs < 0 {
				firstPosAbs = v
			}
			count += v
		}
	}
	if (k - m) % 2 == 0 || m >= k{
		return count
	}
	var minAbs int
	if firstPosAbs < 0 || (lastNegAbs >= 0 && lastNegAbs < firstPosAbs) {
		minAbs = lastNegAbs
	} else {
		minAbs = firstPosAbs
	}
	return count - 2 * minAbs
}