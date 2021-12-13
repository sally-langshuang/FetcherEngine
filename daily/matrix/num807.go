package matrix
//输入：grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
//输出：35

func maxIncreaseKeepingSkyline(grid [][]int) int {
	max := [2][]int{make([]int, len(grid)), make([]int, len(grid[0]))}
	for i:= range grid {
		for j:=range grid[i] {
			if max[0][i] < grid[i][j] {
				max[0][i] = grid[i][j]
			}
			if max[1][j] < grid[i][j] {
				max[1][j] = grid[i][j]
			}
		}
	}
	ans := 0
	for i:= range grid {
		for j:=range grid[i] {
			if max[0][i] <  max[1][j] {
				ans +=  max[0][i]- grid[i][j]
			}else {
				ans += max[1][j] - grid[i][j]
			}
		}
	}
	return ans
}
