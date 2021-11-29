package leet

import "fmt"

func removeElement(nums []int, val int) int {
	j:=0
	for i:=0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	fmt.Println(nums)
	return j
}
