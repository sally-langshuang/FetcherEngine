package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := range obstacleGrid {
		for j := range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
				continue
			}
			if i > 0 {
				obstacleGrid[i][j] += obstacleGrid[i-1][j]
			}
			if j > 0 {
				obstacleGrid[i][j] += obstacleGrid[i][j-1]
			}

		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
