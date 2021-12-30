package dp

func minimumTotal(triangle [][]int) (total int) {
	for i := 1; i < len(triangle); i++ {
		for j := range triangle[i] {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}
	total = triangle[len(triangle)-1][0]
	for _, x := range triangle[len(triangle)-1][1:] {
		if x < total {
			total = x
		}
	}
	return total
}
