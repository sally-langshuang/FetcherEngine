package num_arr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNum506(t *testing.T)  {
	datas := []struct {
		score[]int
		ans []string
	}{
		{[]int{5,4,3,2,1}, []string{"Gold Medal","Silver Medal","Bronze Medal","4","5"}},
		{[]int{10,3,8,9,4}, []string{"Gold Medal","5","Bronze Medal","Silver Medal","4"}},
	}
	for _, x:= range datas {
		fmt.Println(reflect.DeepEqual(findRelativeRanks(x.score), x.ans))
	}

}
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
