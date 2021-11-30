package leet

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func mergeRange(s1, e1, s2, e2 int) []int {
	ans := []int{-1, -1}
	if s1 != -1 || s2 != -1 {
		if x := min(s1, s2); x != -1 {
			ans[0] = x
		} else {
			ans[0] = max(s1, s2)
		}
	}
	ans[1] = max(e1, e2)
	return ans
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if len(nums) == 0 {
		return ans
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		ans = []int{mid, mid}
	}
	left := searchRange(nums[:mid], target)
	if left[0] != -1 {
		ans = mergeRange(ans[0], ans[1], left[0], left[1])
	}
	right := searchRange(nums[mid+1:], target)
	if right[0] != -1 {
		ans = mergeRange(ans[0], ans[1], mid + 1 + right[0], mid + 1 + right[1])
	}
	return ans
}
