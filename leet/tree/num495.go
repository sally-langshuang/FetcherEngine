package tree

func findPoisonedDuration(timeSeries []int, duration int) int {
	lastBegin :=  -duration
	res := 0
	for _, b := range timeSeries {
		if b - lastBegin < duration {
			res += b -lastBegin
		} else {
			res += duration
		}
		lastBegin = b
	}
	return res
}
