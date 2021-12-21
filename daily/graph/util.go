package graph

func newQueue(nums *[]int, less func(i, j int) bool) {
	arr := make([]int, 0)
	for _, i := range *nums {
		Push(&arr, i, less)
	}
	*nums = arr
}

func Push(arr *[]int, val int, less func(i, j int) bool) {
	*arr = append(*arr, val)
	up(arr, len(*arr)-1, less)
}

func Pop(arr *[]int, less func(i, j int) bool) int {
	res := (*arr)[0]
	(*arr)[0] = (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	down(arr, 0, less)
	return res
}

func up(arr *[]int, i int, less func(i, j int) bool) {
	for x := i; x > 0; {
		j := (x - 1) / 2
		if less((*arr)[j], (*arr)[x]) {
			break
		}
		(*arr)[j], (*arr)[x] = (*arr)[x], (*arr)[j]
		x = j
	}
}

func down(arr *[]int, i0 int, less func(i, j int) bool) bool {
	i := i0
	for i < len(*arr) {
		j := i*2 + 1
		if j >= len(*arr) {
			break
		}
		if j2 := j + 1; j2 < len(*arr) && less((*arr)[j2], (*arr)[j]) {
			j = j2
		}
		if less((*arr)[i], (*arr)[j]) {
			break
		}
		(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
		i = j
	}
	return i > i0
}
