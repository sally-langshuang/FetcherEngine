package leet

func search1(nums []int, target int) int {
	k := findK(nums)
	return searchf(target, 0, len(nums), k, nums)
}

func searchf(target, start, end , k int, nums []int) int {
	if end <= start {
		return -1
	}
	mid := (end - start) / 2 + start
	if g(mid, k, len(nums)) == 2 {
		//fmt.Println("why")
	}
	if nums[g(mid, k, len(nums))] == target {
		return g(mid, k, len(nums))
	}
	if x := searchf(target, start, mid, k, nums); x != -1 {
		return x
	}
	if x := searchf(target, mid + 1, end, k, nums); x != -1 {
		return x
	}
	return -1
}

func findK(nums []int) int {
	for i:=0; i < len(nums) - 1; i++{
		if nums[i] > nums[i+1] {
			return len(nums) - i - 1
		}
	}
	return 0
}

func g(i, k, n int)  (j int){
	if i < k {
		return n - k + i
	}
	return i - k
}

func search0(nums []int, target int) int {
	k := findK(nums)
	return searchf(target, 0, len(nums), k, nums)
}

func search(nums []int, target int) int {
	return searchF(0, len(nums), target, nums)
}

func searchF(start, end, target int, nums []int) int{
	if end <= start {
		return -1
	}
	mid := start + (end - start) / 2
	if nums[mid] == target {
		return mid
	}
	if isInLeft(start, mid, end, target, nums){
		return searchF(start, mid, target, nums)
	} else {
		return searchF(mid + 1, end, target, nums)
	}
}

func isInLeft(start, mid, end, target int, nums []int) bool {
	if mid > start {
		if nums[start] <= nums[mid - 1] {
			return target <= nums[mid - 1] && target >= nums[start]
		} else {
			return target > nums[end - 1] || target < nums[mid + 1]
		}
	}
	return false
}
