package num_arr

import (
	"fmt"
	"reflect"
	"testing"
)


func TestNum1005(t *testing.T)  {
	datas := []struct{
		nums []int
		k, ans int
	}{
		{[]int{4, 2, 3}, 1, 5},
		{[]int{3,-1,0,2}, 3, 6},
		{[]int{2,-3,-1,5,-4}, 2, 13},
		{[]int{-4,-2,-3}, 4, 5},
		{[]int{8,-7,-3,-9,1,9,-6,-9,3}, 8, 53},
		{[]int{4,-5,4,-5,9,4,5}, 1, 26},
	}
	for i := range datas {
		if largestSumAfterKNegations(datas[i].nums, datas[i].k) != datas[i].ans {
			fmt.Println(largestSumAfterKNegations(datas[i].nums, datas[i].k), "!=", datas[i].ans )
		}

	}
}

func TestNum35(t *testing.T)  {
	datas := []struct{
		nums []int
		target int
		ans int
	} {
		//{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{1,3,5,6}, 0, 0},
		{[]int{1}, 0, 0},
	}
	for i := range datas {
		fmt.Println(searchInsert(datas[i].nums, datas[i].target) == datas[i].ans)
	}
}

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
