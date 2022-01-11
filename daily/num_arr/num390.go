package num_arr

import "fmt"

func lastRemaining(n int) int {
	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	for len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			arr = append(arr[:i], arr[i+1:]...)
		}
		if len(arr) == 1 {
			break
		}
		fmt.Println(arr)
		for i := len(arr) - 1; i >= 0; i -= 2 {
			arr = append(arr[:i], arr[i+1:]...)
		}
		fmt.Println(arr)
	}
	return arr[0]
}
