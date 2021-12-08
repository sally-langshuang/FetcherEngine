package num_arr

func trap(height []int) int {
	var ans, lastHeight int
	for i, j := 0, len(height) - 1; i < j; {
		if h:= min(height[i], height[j]); h> lastHeight {
			ans += (h - lastHeight) * (j - i - 1) - lastHeight
			lastHeight = h
		} else {
			ans -= h
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
