package num_arr

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	ans := make([]string, len(score))
	arr := make([]Person, len(score))
	for i := range score {
		arr[i] = Person{score: score[i], idx: i}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score >= arr[j].score
	})
	//arr = merge(arr)
	x := []string{"Gold Medal","Silver Medal","Bronze Medal"}
	for i := range arr {
		var val string
		if i < 3 {
			val = x[i]
		} else {
			val = strconv.Itoa(i + 1)
		}
		ans[arr[i].idx] = val
	}
	return ans
}

type Person struct {
	score, idx int
}

func merge(arr []Person)  []Person{
	if len(arr) <= 1 {
		return arr
	}
	ans := make([]Person, len(arr))
	mid := len(arr) / 2
	left := merge(arr[:mid])
	right := merge(arr[mid:])
	i, j := 0, 0
	for ; i < len(left) && j < len(right); {
		if left[i].score >= right[j].score {
			ans[i+j] = left[i]
			i++
			continue
		} else {
			ans[i+j] = right[j]
			j++
			continue
		}
	}
	if i < len(left) {
		for ;i + j < len(arr); i++ {
			ans[i + j] = left[i]
		}
	}
	if j < len(right) {
		for ;i + j < len(arr); j++ {
			ans[i + j] = right[j]
		}
	}
	return ans
}
