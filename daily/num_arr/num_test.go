package num_arr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum786(t *testing.T)  {
	datas := []struct{
		arr, ans []int
		k int
	}{
		{[]int{1, 2, 3, 5}, []int{2, 5}, 3},
		{[]int{1, 7}, []int{1, 7}, 1},
		{[]int{1,2,11,37,83,89}, []int{11, 37}, 11},
	}
	for _, x := range datas {
		fmt.Println(reflect.DeepEqual(kthSmallestPrimeFraction(x.arr, x.k), x.ans))
	}
}
