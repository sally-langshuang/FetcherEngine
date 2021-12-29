package dp

func minFallingPathSum0(matrix [][]int) (total int) {
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[i] {
			pres := []int{matrix[i-1][j]}
			if j > 0 {
				pres = append(pres, matrix[i-1][j-1])
			}
			if j < len(matrix[i])-1 {
				pres = append(pres, matrix[i-1][j+1])
			}
			matrix[i][j] += min(pres...)
		}
	}
	total = matrix[len(matrix)-1][0]
	for _, x := range matrix[len(matrix)-1][1:] {
		if x < total {
			total = x
		}
	}
	return total
}
