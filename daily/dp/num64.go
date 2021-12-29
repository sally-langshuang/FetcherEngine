package dp

func minPathSum(grid [][]int) int {
	for i := range grid {
		for j := range grid[i] {
			vals := []int{}
			if i > 0 {
				vals = append(vals, grid[i-1][j])
			}
			if j > 0 {
				vals = append(vals, grid[i][j-1])
			}
			grid[i][j] += min(vals...)
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	for _, x := range nums[1:] {
		if x < ans {
			ans = x
		}
	}
	return ans
}
