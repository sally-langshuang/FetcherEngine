package leet

import "sort"

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
func find0(sum, count int, nums []int) [][]int {
	//fmt.Printf("find %d, %d, %v\n", sum, count,  nums)
	ans := make([][]int, 0)
	if len(nums) < count || len(nums) == 0  {
		return [][]int{}
	}
	if count == 1 || len(nums) == 1{
		for _, x := range nums{
			if x == sum {
				if len(ans) == 0 {
					ans = append(ans, []int{x})
				}
			}
		}
	} else {
		for _, x:= range find(sum, count, nums[:len(nums) - 1]) {
			if len(x) > 0 && x[len(x) - 1] != nums[len(nums) - 1]{
				ans = append(ans, x)
			}
		}
		for _, x := range find(sum - nums[len(nums) -1], count - 1, nums[:len(nums) - 1]) {
			ans = append(ans, append(x, nums[len(nums) -1],))
		}
	}

	return ans
}

func find(sum, count int, nums []int) [][]int {
	ans := make([][]int, 0)
	for a:= 0; a < len(nums) - 2; a++{
		if a > 0 && nums[a - 1] == nums[a] {
			continue
		}
		for b:= a + 1; b< len(nums) -1; b++ {
			if b > a + 1 && nums[b - 1] == nums[b] {
				continue
			}
			for c:=len(nums) -1; c > b; c-- {
				if c - 1 > b && nums[c - 1] == nums[c] {
					continue
				}
				if nums[a]+nums[b]+nums[c] == 0 {
					ans= append(ans, []int{nums[a], nums[b], nums[c] })
				}
			}
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for a:= 0; a < len(nums) - 2; a++{
		if a > 0 && nums[a - 1] == nums[a] {
			continue
		}
		for b, c:= a + 1, len(nums) -1; b< len(nums) -1; b++ {
			if b > a + 1 && nums[b - 1] == nums[b] {
				continue
			}
			for ; nums[a]+nums[b]+nums[c] > 0 && c > b; c-- {
				if c-1 > b && nums[c-1] == nums[c] {
					continue
				}
			}
			if c == b {
				break
			}
			if nums[a]+nums[b]+nums[c] == 0 {
				ans= append(ans, []int{nums[a], nums[b], nums[c] })
			}
		}
	}
	return ans
}
func threeSumOfficial(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
