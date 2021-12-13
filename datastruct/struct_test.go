package datastruct

import (
	"fmt"
	"sort"
	"testing"
)

func TestStruct(t *testing.T)  {
	nums := []int{1, 3, 5, 2, 5, 6, 1, 9}
	nums2 := map[int]struct{}{}
	nums3 := make([]int,0, len(nums2))
	ans := make([]int, len(nums))
	for i := range nums {
		nums2[nums[i]] = struct{}{}
	}
	for k, _ := range nums2 {
		nums3 = append(nums3, k)
	}
	sort.Ints(nums3)
	c := make([]int, len(nums3))
	for i:= len(nums) -1; i >= 0; i-- {
		n :=  nums[i]
		idx := sort.SearchInts(nums3, n)
		update(c, idx + 1)
		ans[i] = sum(c, idx)
	}
	fmt.Println( nums)
	fmt.Println( nums3)
	fmt.Println( c)
	fmt.Println( ans)
}
func update(c []int, idx int)  {
	for i:= idx; i <= len(c); i = i + i & -i{
		c[i - 1] += 1
	}
}

func sum(c []int, idx int) int {
	ans := 0
	for i:= idx; i > 0; i = i - i&-i {
		ans += c[i - 1]
	}
	return ans
}
