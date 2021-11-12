package dp

func maxArea(height []int) int {
	maxS := 0
	min := func(a, b int) int {if b < a {return b}; return a}
	for i, j:= 0, len(height) - 1; i < j; {
		s := (j - i) * min(height[i], height[j])
		if s > maxS {
			maxS = s
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxS
}
