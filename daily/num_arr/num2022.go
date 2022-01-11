package num_arr

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = original[i*n : (i+1)*n]
	}
	return ans
}
