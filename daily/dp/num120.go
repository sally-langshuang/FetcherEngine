package dp

func minimumTotal(triangle [][]int) (total int) {
	total = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		minVal := triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			if triangle[i][j] < minVal {
				minVal = triangle[i][j]
			}
		}
		total += minVal
	}
	return total
}
