package leet

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if i > 0 && nums[j] == nums[i - 1] {
			continue
		}
		if i != j {
			nums[i] = nums[j]
		}
		i++
	}
	return i
}
