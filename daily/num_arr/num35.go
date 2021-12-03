package num_arr

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return searchInsert(nums[:mid], target)
	}
	return mid + 1 + searchInsert(nums[mid + 1:], target)
}