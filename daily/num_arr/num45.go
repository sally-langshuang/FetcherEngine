package num_arr
// todo
func jump(nums []int) int {
	minStep, stat := 0, map[int]int{0: 0}
	for i := 1; i < len(nums); i++ {
		if v, ok := stat[minStep + 1]; ok && v < nums[i-1] || !ok {
			stat[minStep + 1] = nums[i-1]
		}
		intSize := 32 << (^uint(0) >> 63) // 32 or 64
		minI, tmp := 1<<(intSize-1) - 1, map[int]int{}
		for step, residue := range stat {
			if residue >= 1 {
				tmp[step] = residue - 1
				if step <  minI {
					minI = step
				}
			}
		}
		minStep = minI
		stat = tmp
		//fmt.Println(i, tmp, minStep)
	}
	return minStep
}


