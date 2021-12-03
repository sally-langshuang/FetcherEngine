package num_arr

func heakInit(nums []int) {
	for i:= (len(nums) / 2 - 1); i >= 0; i-- {
		down(i, len(nums), nums)
	}
}

func heakSort(nums []int) []int {
	ans := make([]int, len(nums))
	for i:= (len(nums) / 2 - 1); i >= 0; i-- {
		down(i, len(nums), nums)
	}
	for i := range nums {
		ans[i] = pop(&nums)
	}
	return ans
}


func pop(nums *[]int) int {
	v := (*nums)[0]
	n := len(*nums) - 1
	swap(0, n, *nums)
	down(0, n, *nums)
	*nums = (*nums)[: n]
	return v
}

func push(v int, nums *[]int)  {
	*nums = append(*nums, v)
	up(len(*nums) - 1, *nums)
}

func swap(i, j int, nums []int)  {
	nums[i], nums[j] = nums[j], nums [i]
}

func upper(i, j int, nums []int) bool {
	return nums[i] < nums[j]
}

func up(i int, nums []int)  {
	for {
		j := (i - 1) / 2
		if j == i || upper(j, i, nums) {
			break
		}
		swap(i, j, nums)
		i = j
	}
}

func down(i, n int, nums []int) {
	for {
		j := 2 * i + 1
		if j >= n {
			return
		}
		if right := j + 1; right < n && upper(right, j, nums) {
			j = right
		}
		if upper(i, j, nums) {
			break
		}
		swap(i, j, nums)
		i = j
	}
}
