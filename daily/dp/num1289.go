package dp

func minFallingPathSum(grid [][]int) (minPath int) {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	minFirst, minFirstVal, minSecondVal := findMinFirstSecondIdx(grid[0])
	for i := 1; i < n; i++ {
		for j := range grid[i] {
			preMin := minFirstVal
			if j == minFirst {
				preMin = minSecondVal
			}
			grid[i][j] += preMin
		}
		minFirst, minFirstVal, minSecondVal = findMinFirstSecondIdx(grid[i])
	}
	return minFirstVal
}

func findMinFirstSecondIdx(arr []int) (minFirst int, minFirstVal int, minSecondVal int) {
	var minSecond int
	for j := range arr[1:] {
		if arr[j] < arr[minFirst] {
			minFirst = j
		}
	}
	if minFirst == 0 {
		minSecond = 1
	}
	for j := range arr {
		if j == minFirst {
			continue
		}
		if arr[j] < arr[minSecond] {
			minSecond = j
		}
	}
	return minFirst, arr[minFirst], arr[minSecond]
}
