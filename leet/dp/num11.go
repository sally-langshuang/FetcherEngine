package dp
//若向内 移动短板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
//若向内 移动长板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j])​ 不变或变小，因此下个水槽的面积 一定变小 。

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
