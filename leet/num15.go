package leet
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func sortByRange(head, tail int, nums *[]int) {
	if head == tail {
		return
	}
	h, t := head, tail
	for ; head != tail; {
		for (*nums)[h] <= (*nums)[t] && h <= t {
			if h == t {
				goto stop
			}
			t--
		}
		temp := (*nums)[t]
		(*nums)[t] = (*nums)[h]
		(*nums)[h] = temp
		for (*nums)[h] <= (*nums)[t] && h <= t {
			if h == t {
				goto stop
			}
			h++
		}
		temp = (*nums)[t]
		(*nums)[t] = (*nums)[h]
		(*nums)[h] = temp
	}
stop:
	if h-1 > head {
		sortByRange(head, h-1, nums)
	}
	if t+1 < tail {
		sortByRange(t+1, tail, nums)
	}

}
func Sort(nums[]int) []int {
	sortByRange(0, len(nums) - 1, &nums)
	return nums
}
func find(sum, count int, nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < count || len(nums) == 0  {
		return [][]int{}
	}
	if len(nums) > 1 && nums[len(nums) -1] == nums[len(nums) -2] {
		return find(sum, count, nums[:len(nums) - 1])
	}
	if count == 1 || len(nums) == 1{
		for _, x := range nums{
			if x == sum {
				ans = append(ans, []int{x})
			}
		}
	} else {
		ans = append(ans, find(sum, count, nums[:len(nums) - 1])...)
		for _, x := range find(sum - nums[len(nums) -1], count - 1, nums[:len(nums) - 1]) {
			ans = append(ans, append(x, nums[len(nums) -1],))
		}
	}

	return ans
}
func threeSum(nums []int) [][]int {
	Sort(nums)
	return find(0, 3, nums)

}