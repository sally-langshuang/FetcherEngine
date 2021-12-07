package matrix

var steps = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isInRange(grid[][]int, now []int) bool {
	return now[0] >= 0 && now[1] >= 0 && now[0] <len(grid) && now[1] < len(grid[0])
}

func nextSteps(grid [][]int, now []int, filter func([]int) bool) [][]int {
	if !isInRange(grid, now) {return [][]int{}}
	ans := [][]int{}
	for i := range steps {
		next := []int{now[0] + steps[i][0], now[1] + steps[i][1]}
		//fmt.Println("-->",next)
		if filter(next){
			ans = append(ans, next)
		}
	}
	return ans
}

func isTodo(done, grid [][]int, now []int, same int) bool {
	return isInRange(done, now) && done[now[0]][now[1]] == 0 && grid[now[0]][now[1]] == same
}

func doColorBorder(grid, done [][]int, now []int, color int, same int) {
	if !isInRange(done, now) || grid[now[0]][now[1]] != same || done[now[0]][now[1]] != 0{
		return
	}
	done[now[0]][now[1]] = 1

	if nexts := nextSteps(grid, now, func(n []int) bool{
		return !isInRange(grid, n) || (done[n[0]][n[1]] == 0 && grid[n[0]][n[1]] != same)
	}); len(nexts) != 0 {
		grid[now[0]][now[1]] = color

	}
	nexts := nextSteps(grid, now, func(next []int) bool {
		return isInRange(grid, next) && grid[next[0]][next[1]] == same
	})
	for i:= range nexts {
		doColorBorder(grid, done, nexts[i], color, same)
	}

}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	done := make([][]int, len(grid))
	for i:= range done {
		done[i] = make([]int, len(grid[i]))
	}
	doColorBorder(grid, done, []int{row, col}, color, grid[row][col])
	return grid
}