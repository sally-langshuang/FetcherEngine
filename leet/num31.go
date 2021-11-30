package leet

//输入：
//[1,3,2]
//输出：
//[3,1,2]
//预期结果：
//[2,1,3]

func insertIdx(num int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mid:= len(nums) / 2
	if nums[mid] == num {
		return mid + 1
	} else if num < nums[mid] {
		return insertIdx(num, nums[:mid])
	} else {
		return mid  + 1 + insertIdx(num, nums[mid + 1:])
	}
}

func insertHead(nums []int) {
	if len(nums) <= 1 {
		return
	}
	head := nums[0]
	i := insertIdx(nums[0], nums[1:])
	for j:=0; j < i; j++ {
		nums[j] = nums[j+1]
	}
	nums[i] = head
}
func reverse(nums []int)  {
	for i:=0 ; i < len(nums) / 2;i ++ {
		nums[i], nums[len(nums) - 1 - i] = nums[len(nums) - 1 - i], nums[i]
	}

}
//[1,3,2]
//预期结果：
//[2,1,3]
func nextPermutation(nums []int) {
	for i:= len(nums) - 1; i >= 0; i-- {
	   if i == 0 || nums[i - 1] < nums[i] {
		   reverse(nums[i:])
		   if i != 0 {
			   for j:=i; j < len(nums); j++ {
				   if nums[j] > nums[i - 1] {
					   nums[i - 1], nums[j] = nums[j], nums[i - 1]
					   return
				   }
			   }
		   }
		   return
	   }
	}
	reverse(nums)
}